package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"graph-view-project/models"
	"log"
	"os"

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
	query := `SELECT 1 FROM users WHERE login = $1`
	var email, password string
	err := db.Connection.QueryRow(query, login).Scan(email, password)
	if err != nil {
		return nil, err
	}
	return &models.User{Id: 0, Login: login, Email: email, Password: password}, nil
}

func (db *PostgresDB) GetUser(id int) (*models.User, error) {
	query := `SELECT 1 FROM users WHERE id = $1`
	var login, email, password string
	err := db.Connection.QueryRow(query, id).Scan(login, email, password)
	if err != nil {
		return nil, err
	}
	return &models.User{Id: id, Login: login, Email: email, Password: password}, nil
}

func (db *PostgresDB) CreateGraph(user models.User, graph models.Graph) (*models.Graph, error) {
	query := `INSERT INTO graphs (name, description) VALUES ($1, $2) returning id`
	err := db.Connection.QueryRow(query, graph.Name, graph.Description).Scan(&graph.Id)
	if err != nil {
		return nil, err
	}
	err = db.UpdateUser(user)
	if err != nil {
		revert := `DELETE FROM graphs WHERE id = $1`
		db.Connection.QueryRow(revert, graph.Id)
		return nil, err
	}
	return &graph, nil
}
func (db *PostgresDB) CreateNode(graph models.Graph, node models.Node) (*models.Node, error) {
	return nil, nil
}
func (db *PostgresDB) CreateEdge(graph models.Graph, edge models.Edge) (*models.Edge, error) {
	return nil, nil
}

func (db *PostgresDB) GetGraph(id int) (*models.Graph, error) {
	return nil, error(nil)
}
func (db *PostgresDB) GetEdge(id int) (*models.Edge, error) {
	return nil, error(nil)
}
func (db *PostgresDB) GetNode(id int) (*models.Node, error) {
	return nil, error(nil)
}

func (db *PostgresDB) UpdateUser(newUser models.User) error {
	query := `UPDATE users set login = $1, email = $2 , password = $3 , graphs = $4 WHERE id = $5`
	graphIds := make([]int, 0)
	for _, g := range newUser.Graphs {
		graphIds = append(graphIds, g.Id)
	}
	err := db.Connection.QueryRow(query, newUser.Login, newUser.Email, newUser.Password, pq.Array(graphIds), newUser.Id).Err()
	return err
}

func (db *PostgresDB) UpdateGraph(newGraph models.Graph) error {
	return error(nil)
}

func (db *PostgresDB) UpdateEdge(newEdge models.Edge) error {
	return error(nil)
}

func (db *PostgresDB) UpdateNode(newNode models.Node) error {
	return error(nil)
}
