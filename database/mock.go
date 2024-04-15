package database

import (
	"errors"
	"graph-view-project/models"
	"graph-view-project/server/utils"
)

type MockDB struct {
	users  map[int]*models.User
	graphs map[int]*models.Graph
	nodes  map[int]*models.Node
	edges  map[int]*models.Edge
}

func (db *MockDB) Connect(databaseName string) error {
	_ = databaseName //this parameter is useless because this method is just a Mock
	db.users = make(map[int]*models.User)
	db.graphs = make(map[int]*models.Graph)
	db.edges = make(map[int]*models.Edge)
	db.nodes = make(map[int]*models.Node)

	hash, _ := utils.GenerateHashPassword("password")
	_, _ = db.CreateUser(models.User{Login: "user", Email: "user@user", Password: hash, Graphs: make([]models.Graph, 0)})

	hash, _ = utils.GenerateHashPassword("12345678password")
	_, _ = db.CreateUser(models.User{Login: "12345", Email: "123456", Password: hash, Graphs: make([]models.Graph, 0)})
	return nil
}

func (db *MockDB) Disconnect() error {
	db.users = nil
	db.graphs = nil
	db.edges = nil
	db.nodes = nil
	return nil
}

func (db *MockDB) CompleteReset() error {
	db.users = make(map[int]*models.User)
	db.graphs = make(map[int]*models.Graph)
	db.edges = make(map[int]*models.Edge)
	db.nodes = make(map[int]*models.Node)
	return nil
}

func (db *MockDB) CreateUser(user models.User) (*models.User, error) {
	if _, err := db.GetUserByLogin(user.Login); err == nil {
		return nil, errors.New("user already exists")
	}

	user.Id = len(db.graphs)
	db.users[user.Id] = &user
	err := db.UpdateUser(user)
	newUser := user
	return &newUser, err
}

func (db *MockDB) CreateGraph(user models.User, graph models.Graph) (*models.Graph, error) {
	graph.Id = len(db.graphs)
	db.graphs[graph.Id] = &graph
	err := db.UpdateUser(user)
	newGraph := graph
	return &newGraph, err
}

func (db *MockDB) CreateNode(graph models.Graph, node models.Node) (*models.Node, error) {
	node.Id = len(db.nodes)
	db.nodes[node.Id] = &node
	err := db.UpdateGraph(graph)
	newNode := node
	return &newNode, err
}

func (db *MockDB) CreateEdge(graph models.Graph, edge models.Edge) (*models.Edge, error) {
	edge.Id = len(db.edges)
	db.edges[edge.Id] = &edge
	err := db.UpdateGraph(graph)
	newEdge := edge
	return &newEdge, err
}

func (db *MockDB) GetUser(id int) (*models.User, error) {
	if len(db.users) <= id {
		return nil, errors.New("user does not exist")
	}

	return db.users[id], nil
}

func (db *MockDB) GetUserByLogin(login string) (*models.User, error) {
	for _, u := range db.users {
		if (*u).Login == login {
			return u, nil
		}
	}

	return nil, errors.New("user does not exist")
}

func (db *MockDB) GetGraph(id int) (*models.Graph, error) {
	if len(db.graphs) <= id {
		return nil, errors.New("graph does not exist")
	}

	return db.graphs[id], nil
}
func (db *MockDB) GetEdge(id int) (*models.Edge, error) {
	edge, ok := db.edges[id]
	if !ok {
		return nil, errors.New("graph does not exist")
	}

	return edge, nil
}
func (db *MockDB) GetNode(id int) (*models.Node, error) {
	node, ok := db.nodes[id]
	if !ok {
		return nil, errors.New("node does not exist")
	}

	return node, nil
}

func (db *MockDB) UpdateUser(newUser models.User) error {
	user, err := db.GetUser(newUser.Id)
	if err != nil {
		return err
	}

	*user = newUser
	return nil
}

func (db *MockDB) UpdateGraph(newGraph models.Graph) error {
	graph, err := db.GetGraph(newGraph.Id)
	if err != nil {
		return err
	}

	*graph = newGraph
	return nil
}

func (db *MockDB) UpdateEdge(newEdge models.Edge) error {
	edge, err := db.GetEdge(newEdge.Id)
	if err != nil {
		return err
	}

	*edge = newEdge
	return nil
}

func (db *MockDB) UpdateNode(newNode models.Node) error {
	node, err := db.GetNode(newNode.Id)
	if err != nil {
		return err
	}

	*node = newNode
	return nil
}
