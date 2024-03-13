package database

import (
	"errors"
	"graph-view-project/models"
	"graph-view-project/server/utils"
)

type MockDB struct {
	users  map[string]*models.User
	graphs map[int]*models.Graph
	nodes  map[int]*models.Node
	edges  map[int]*models.Edge
}

func (db *MockDB) Connect(databaseName string) error {
	_ = databaseName //this parameter is useless because this method is just a Mock
	db.users = make(map[string]*models.User)

	hash, _ := utils.GenerateHashPassword("password")
	_ = db.CreateUser(models.User{Id: 1, Login: "user", Email: "user@user", Password: hash, Graphs: make([]models.Graph, 0)})

	hash, _ = utils.GenerateHashPassword("12345678password")
	_ = db.CreateUser(models.User{Id: 2, Login: "12345", Email: "123456", Password: hash, Graphs: make([]models.Graph, 0)})
	return nil
}

func (db *MockDB) Disconnect() error {
	db.users = nil
	db.graphs = nil
	return nil
}

func (db *MockDB) CreateUser(user models.User) error {
	_, ok := db.users[user.Login]
	if ok {
		return errors.New("user already exists")
	}

	db.users[user.Login] = &user

	return nil
}

func (db *MockDB) CreateGraph(user models.User, graph models.Graph) error {
	_ = user //parameter user is needed to add a link to this new graph into user data.
	//therefore, here in Mock we use pointers to objects, so we do not have to change them manually
	_, ok := db.graphs[graph.GetId()]
	if ok {
		return errors.New("graph already exists")
	}

	db.graphs[graph.GetId()] = &graph

	return nil
}
func (db *MockDB) CreateNode(graph models.Graph, node models.Node) error {
	_ = graph //in case of questions, read comment in the CreateGraph func
	_, ok := db.nodes[node.GetId()]
	if ok {
		return errors.New("graph already exists")
	}

	db.nodes[node.GetId()] = &node

	return nil
}
func (db *MockDB) CreateEdge(graph models.Graph, edge models.Edge) error {
	_ = graph //in case of questions, read comment in the CreateGraph func
	_, ok := db.nodes[edge.GetId()]
	if ok {
		return errors.New("graph already exists")
	}

	db.edges[edge.GetId()] = &edge

	return nil
}

func (db *MockDB) GetUserByLogin(login string) (*models.User, error) {
	user, ok := db.users[login]
	if !ok {
		return nil, errors.New("user does not exist")
	}

	return user, nil
}

func (db *MockDB) GetGraphById(id int) (*models.Graph, error) {
	graph, ok := db.graphs[id]
	if !ok {
		return nil, errors.New("graph does not exist")
	}

	return graph, nil
}
func (db *MockDB) GetEdgeById(id int) (*models.Edge, error) {
	edge, ok := db.edges[id]
	if !ok {
		return nil, errors.New("graph does not exist")
	}

	return edge, nil
}
func (db *MockDB) GetNodeById(id int) (*models.Node, error) {
	node, ok := db.nodes[id]
	if !ok {
		return nil, errors.New("node does not exist")
	}

	return node, nil
}
