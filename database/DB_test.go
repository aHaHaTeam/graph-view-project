package database

import (
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"graph-view-project/models"
	"graph-view-project/server/utils"
	"os"
	"testing"
)

var testedDB DataBase

func Setup() {
	testedDB = &MockDB{}
	_ = testedDB.Connect("test")
}

func TearDown() {
	_ = testedDB.Disconnect()
}

func CleanUp() {
	TearDown()
	Setup()
}

func TestConnectDisconnect(t *testing.T) {
	t.Cleanup(CleanUp)
	err := testedDB.Connect("test")
	assert.Nil(t, err)

	err = testedDB.Disconnect()
	assert.Nil(t, err)
}

func TestCreateUser(t *testing.T) {
	t.Cleanup(CleanUp)

	hash, _ := utils.GenerateHashPassword("PaSsWoRd")
	user := models.User{Id: 1, Login: "user2", Email: "user@user", Password: hash, Graphs: make([]models.Graph, 0)}
	err := testedDB.CreateUser(user)
	assert.Nil(t, err)
	err = testedDB.CreateUser(user)
	assert.NotNil(t, err)
}

func TestCreateGraph(t *testing.T) {
	t.Cleanup(CleanUp)
	user := models.User{}
	_ = testedDB.CreateUser(user)

	graph := models.Graph{Id: 1, Edges: make([]models.Edge, 0), Nodes: make([]models.Node, 0)}

	err := testedDB.CreateGraph(user, graph)
	assert.Nil(t, err)

	err = testedDB.CreateGraph(user, graph)
	assert.NotNil(t, err)
}

func TestCreateNode(t *testing.T) {
	t.Cleanup(CleanUp)
	user := models.User{}
	_ = testedDB.CreateUser(user)

	graph := models.Graph{}
	_ = testedDB.CreateGraph(user, graph)

	node := models.Node{Id: 1}

	err := testedDB.CreateNode(graph, node)
	assert.Nil(t, err)

	err = testedDB.CreateNode(graph, node)
	assert.NotNil(t, err)
}

func TestCreateEdge(t *testing.T) {
	t.Cleanup(CleanUp)

	user := models.User{}
	_ = testedDB.CreateUser(user)

	graph := models.Graph{}
	_ = testedDB.CreateGraph(user, graph)

	edge := models.Edge{Id: 1}

	err := testedDB.CreateEdge(graph, edge)
	assert.Nil(t, err)

	err = testedDB.CreateEdge(graph, edge)
	assert.NotNil(t, err)
}

func TestUpdateUser(t *testing.T) {
	t.Cleanup(CleanUp)

	user := models.User{
		Id:       1,
		Login:    "1",
		Email:    "1",
		Password: "1",
		Graphs:   make([]models.Graph, 0),
	}
	_ = testedDB.CreateUser(user)

	graph := models.Graph{Id: 239, Nodes: nil, Edges: nil}
	user.Graphs = append(user.Graphs, graph)

	userInDb, _ := testedDB.GetUserByLogin("1")
	assert.Equal(t, 0, len(userInDb.Graphs))

	err := testedDB.UpdateUserByLogin(user)
	assert.Nil(t, err)

	userInDb, _ = testedDB.GetUserByLogin("1")
	assert.Equal(t, 1, len(userInDb.Graphs))

	user.Login = "invalidLogin"
	err = testedDB.UpdateUserByLogin(user)
	assert.NotNil(t, err)
}

func TestUpdateGraph(t *testing.T) {
	t.Cleanup(CleanUp)

	user := models.User{}
	_ = testedDB.CreateUser(user)
	graph := models.Graph{
		Id:    1,
		Nodes: make([]models.Node, 0),
		Edges: make([]models.Edge, 0),
	}
	_ = testedDB.CreateGraph(user, graph)

	node := models.Node{}
	edge := models.Edge{}
	graph.Nodes = append(graph.Nodes, node)
	graph.Edges = append(graph.Edges, edge)

	graphInDB, _ := testedDB.GetGraphById(1)
	assert.Equal(t, 0, len(graphInDB.Nodes))

	_ = testedDB.UpdateGraphById(graph.Id, graph)
	graphInDB, _ = testedDB.GetGraphById(1)
	assert.Equal(t, 1, len(graphInDB.Nodes))
	assert.Equal(t, edge, graphInDB.Edges[0])
}

func TestMain(m *testing.M) {
	Setup()
	code := m.Run()
	TearDown()
	os.Exit(code)
}
