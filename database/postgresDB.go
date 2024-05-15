package database

import (
	"database/sql"
	"errors"
	"fmt"
	"graph-view-project/models"
	"log"
	"os"

	"github.com/lib/pq"

	"github.com/joho/godotenv"
)

type PostgresDB struct {
	Connection *sql.DB
}

func (db *PostgresDB) Connect(databaseName string) error {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	_ = databaseName // deprecated parameter
	user := os.Getenv("DB1_USER")
	password := os.Getenv("DB1_PASSWORD")
	host := os.Getenv("DB1_HOST")
	port := os.Getenv("DB1_PORT")
	dbName := os.Getenv("DB1_DBNAME")
	template := "postgres://%s:%s@%s:%s/%s"

	connStr := fmt.Sprintf(template, user, password, host, port, dbName)

	connection_, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	db.Connection = connection_
	return nil
}

func (db *PostgresDB) Disconnect() error {
	err := db.Connection.Close()
	return err
}

func (db *PostgresDB) CompleteReset() error {
	_, err := db.Connection.Exec(`CREATE TABLE IF NOT EXISTS users
(
    id       SERIAL PRIMARY KEY,
    login    VARCHAR(239),
    email    VARCHAR(239),
    password VARCHAR(239),
    graphs   integer[]
);

CREATE TABLE IF NOT EXISTS graphs
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(239),
    description VARCHAR(239),
    nodes       INTEGER[],
    edges       INTEGER[],

    defaultNodeSize  FLOAT DEFAULT 1,
    DefaultNodeColor INTEGER DEFAULT 0,
    DefaultNodeShape INTEGER DEFAULT 0,

    DefaultEdgeWidth FLOAT DEFAULT 1,
    DefaultEdgeColor INTEGER DEFAULT 0,
    DefaultEdgeShape INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS nodes
(
    id    SERIAL PRIMARY KEY,
    edges INTEGER[],

    name  VARCHAR(239),
    data  bytea,
    size  FLOAT DEFAULT 1,
    color INTEGER DEFAULT 0,
    shape INTEGER DEFAULT 0

);

CREATE TABLE IF NOT EXISTS edges
(
    id          SERIAL PRIMARY KEY,
    begin       INTEGER,
    "end"       INTEGER,

    name        VARCHAR(239) DEFAULT 'not specified',
    description VARCHAR(239) DEFAULT 'not null',

    width       FLOAT DEFAULT 1.0,
    color       INTEGER DEFAULT 0,
    shape       INTEGER DEFAULT 0
);

DELETE FROM users where id < 10;
DELETE FROM graphs where id < 30;
DELETE FROM nodes where id < 10;
DELETE FROM edges where id < 10;

TRUNCATE TABLE users RESTART IDENTITY;
TRUNCATE TABLE graphs RESTART IDENTITY;
TRUNCATE TABLE nodes RESTART IDENTITY;
TRUNCATE TABLE edges RESTART IDENTITY;`)

	return err
}

func (db *PostgresDB) CreateUser(user models.User) (*models.User, error) {
	var loginTaken bool
	check := `Select exists (select true from users where login = $1)`
	err := db.Connection.QueryRow(check, user.Login).Scan(&loginTaken)
	if err != nil {
		return nil, err
	}
	if loginTaken {
		return nil, errors.New("user already exists")
	}

	query := `INSERT INTO users (login, email, password) VALUES ($1, $2, $3) returning id`
	err = db.Connection.QueryRow(query, user.Login, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *PostgresDB) GetUserByLogin(login string) (*models.User, error) {
	query := `SELECT id, login, email, password, graphs FROM users WHERE login = $1`
	var email, password string
	var id int
	var graphIds pq.Int64Array
	err := db.Connection.QueryRow(query, login).Scan(&id, &login, &email, &password, &graphIds)

	graphs := make([]*models.Graph, len(graphIds))
	for i, gId := range graphIds {
		graphs[i], err = db.GetGraph(models.GraphId(gId))
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return models.NewUser(models.UserId(id), login, email, password, graphs), nil
}

func (db *PostgresDB) GetUser(id models.UserId) (*models.User, error) {
	query := `SELECT id, login, email, password, graphs FROM users WHERE id = $1`
	var email, password, login string
	var graphIds pq.Int64Array
	err := db.Connection.QueryRow(query, id).Scan(&id, &login, &email, &password, &graphIds)

	graphs := make([]*models.Graph, len(graphIds))
	for i, gId := range graphIds {
		graphs[i], err = db.GetGraph(models.GraphId(gId))
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return models.NewUser(id, login, email, password, graphs), nil
}

func (db *PostgresDB) CreateGraph(user models.User, graph models.Graph) (*models.Graph, error) {
	query := `INSERT INTO graphs (name, description) VALUES ($1, $2) returning id`
	err := db.Connection.QueryRow(query, graph.Name, graph.Description).Scan(&graph.Id)
	if err != nil {
		return nil, err
	}
	err1 := db.UpdateUser(user)
	err2 := db.UpdateGraph(graph)
	if err1 != nil {
		revert := `DELETE FROM graphs WHERE id = $1`
		db.Connection.QueryRow(revert, graph.Id)
		return nil, err1
	}
	if err2 != nil {
		revert := `DELETE FROM graphs WHERE id = $1`
		db.Connection.QueryRow(revert, graph.Id)
		return nil, err2
	}
	return &graph, nil
}
func (db *PostgresDB) CreateNode(graph models.Graph, node models.Node) (*models.Node, error) {
	query := `INSERT INTO nodes (name, size, color, shape) VALUES ($1, $2, $3, $4) returning id`
	err := db.Connection.QueryRow(query, node.Name, node.Size, models.ColorToInt(node.Color), models.NodeShapeToInt(node.Shape)).Scan(&node.Id)
	if err != nil {
		return nil, err
	}
	err = db.UpdateGraph(graph)
	if err != nil {
		revert := `DELETE FROM nodes WHERE id = $1`
		db.Connection.QueryRow(revert, node.Id)
		return nil, err
	}
	return &node, nil
}

func (db *PostgresDB) CreateEdge(graph models.Graph, edge models.Edge) (*models.Edge, error) {
	query := `INSERT INTO edges (begin, "end", name, description, width, color, shape) VALUES ($1, $2, $3, $4, $5, $6, $7) returning id`
	err := db.Connection.QueryRow(query, edge.Begin, edge.End, edge.Name, edge.Description, edge.Width, models.ColorToInt(edge.Color), models.EdgeShapeToInt(edge.Shape)).Scan(&edge.Id)
	if err != nil {
		return nil, err
	}
	err = db.UpdateGraph(graph)
	if err != nil {
		revert := `DELETE FROM edges WHERE id = $1`
		db.Connection.QueryRow(revert, edge.Id)
		return nil, err
	}
	return &edge, nil
}

func (db *PostgresDB) GetGraph(id models.GraphId) (*models.Graph, error) {
	query := `SELECT name, description, nodes, edges, 
       defaultnodesize, defaultnodecolor, defaultedgeshape,
       defaultedgewidth, defaultedgecolor, defaultedgeshape
			FROM graphs WHERE id = $1`
	var name, description string
	var nodeIds, edgeIds pq.Int64Array
	var defaultNodeSize, defaultEdgeWidth float32
	var defaultNodeColor, defaultEdgeColor int
	var defaultNodeShape, defaultEdgeShape int
	err := db.Connection.QueryRow(query, id).Scan(&name, &description, &nodeIds, &edgeIds,
		&defaultNodeSize, &defaultNodeColor, &defaultNodeShape,
		&defaultEdgeWidth, &defaultEdgeColor, &defaultEdgeShape)

	nodes := make([]*models.Node, len(nodeIds))
	edges := make([]*models.Edge, len(edgeIds))
	for i, nId := range nodeIds {
		nodes[i], err = db.GetNode(models.NodeId(nId))
		if err != nil {
			return nil, err
		}
	}

	for i, eId := range edgeIds {
		edges[i], err = db.GetEdge(models.EdgeId(eId))
		if err != nil {
			return nil, err
		}
	}

	return models.NewGraph(id, name, description, nodes, edges,
		defaultNodeSize, models.ColorFromInt(defaultNodeColor), models.NodeShapeFromInt(defaultNodeShape),
		defaultEdgeWidth, models.ColorFromInt(defaultEdgeColor), models.EdgeShapeFromInt(defaultEdgeShape)), nil
}

func (db *PostgresDB) GetEdge(id models.EdgeId) (*models.Edge, error) {
	query := `SELECT begin, "end", name, description, width, color, shape FROM edges WHERE id = $1`

	var begin, end int
	var name, description string
	var width float32
	var color, shape int
	err := db.Connection.QueryRow(query, id).Scan(&begin, &end, &name, &description, &width, &color, &shape)
	if err != nil {
		return nil, err
	}
	return models.NewEdge(id, begin, end, name, description, width, models.ColorFromInt(color), models.EdgeShapeFromInt(shape)), nil
}

func (db *PostgresDB) GetNode(id models.NodeId) (*models.Node, error) {
	query := `SELECT edges, name, data, size, color, shape FROM nodes WHERE id = $1`

	var name string
	var size float32
	var color, shape int
	var data []byte
	var edges []models.NodeId
	err := db.Connection.QueryRow(query, id).Scan(pq.Array(&edges), &name, pq.Array(&data), &size, &color, &shape)
	if err != nil {
		return nil, err
	}
	return models.NewNode(id, edges, name, data, size, models.ColorFromInt(color), models.NodeShapeFromInt(shape)), nil
}

func (db *PostgresDB) UpdateUser(newUser models.User) error {
	_, err := db.GetUser(newUser.Id)
	if err != nil {
		return err
	}
	query := `UPDATE users SET login = $1, email = $2 , password = $3 , graphs = $4 WHERE id = $5`

	graphIds := make([]int, len(newUser.Graphs))
	for i, g := range newUser.Graphs {
		graphIds[i] = int(g.Id)
	}

	err = db.Connection.QueryRow(query, newUser.Login, newUser.Email, newUser.Password, pq.Array(graphIds), newUser.Id).Err()
	return err
}

func (db *PostgresDB) UpdateGraph(newGraph models.Graph) error {
	query := `UPDATE graphs SET name = $1, description = $2, nodes = $3, edges = $4, defaultnodesize = $5, defaultnodecolor = $6, defaultnodeshape = $7, defaultedgewidth = $8, defaultedgecolor = $9, defaultedgeshape = $10 WHERE id = $11`
	nodeIds := make([]int, len(newGraph.Nodes))
	for i, n := range newGraph.Nodes {
		nodeIds[i] = int(n.Id)
	}
	edgeIds := make([]int, len(newGraph.Edges))
	for i, e := range newGraph.Edges {
		edgeIds[i] = int(e.Id)
	}

	err := db.Connection.QueryRow(query, newGraph.Name, newGraph.Description, pq.Array(nodeIds), pq.Array(edgeIds),
		newGraph.DefaultNodeSize, models.ColorToInt(newGraph.DefaultNodeColor), models.NodeShapeToInt(newGraph.DefaultNodeShape),
		newGraph.DefaultEdgeWidth, models.ColorToInt(newGraph.DefaultEdgeColor), models.EdgeShapeToInt(newGraph.DefaultEdgeShape),
		newGraph.Id).Err()
	return err
}

func (db *PostgresDB) UpdateEdge(newEdge models.Edge) error {
	query := `UPDATE edges SET begin = $1, "end" = $2 , name = $3 , description = $4, 
                 width = $5, color = $6, shape = $7 WHERE id = $8`

	err := db.Connection.QueryRow(query, newEdge.Begin, newEdge.End, newEdge.Name, newEdge.Description,
		newEdge.Width, models.ColorToInt(newEdge.Color), models.EdgeShapeToInt(newEdge.Shape)).Err()
	return err
}

func (db *PostgresDB) UpdateNode(newNode models.Node) error {
	query := `UPDATE nodes SET edges = $1, name = $2 , data = $3 , 
                 size = $4, color = $5, shape = $6 WHERE id = $7`

	err := db.Connection.QueryRow(query, pq.Array(newNode.AdjacentNodes), newNode.Name, newNode.Data,
		newNode.Size, models.ColorToInt(newNode.Color), models.NodeShapeToInt(newNode.Shape)).Err()
	return err
}
