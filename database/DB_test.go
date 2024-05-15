package database

import (
	"graph-view-project/models"
	"graph-view-project/server/utils"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var testedDB DataBase

func Setup() {
	testedDB = &PostgresDB{}
	_ = testedDB.Connect("test")
	_ = testedDB.CompleteReset()
}

func TearDown() {
	_ = testedDB.Disconnect()
}

func CleanUp() {
	TearDown()
	Setup()
}

func TestConnectDisconnect(t *testing.T) {
	CleanUp()

	err := testedDB.Connect("test")
	assert.Nil(t, err)

	err = testedDB.Disconnect()
	assert.Nil(t, err)
}

func TestCreateUser(t *testing.T) {
	CleanUp()

	hash, _ := utils.GenerateHashPassword("PaSsWoRd")
	user := models.User{Id: 1, Login: "user2", Email: "user@user", Password: hash, Graphs: make([]*models.Graph, 0)}
	_, err := testedDB.CreateUser(user)
	assert.Nil(t, err)
	_, err = testedDB.CreateUser(user)
	assert.NotNil(t, err)
}

func TestCreateGraph(t *testing.T) {
	CleanUp()

	user := &models.User{Id: 100, Login: "graphOwner", Graphs: make([]*models.Graph, 0)}
	user, _ = testedDB.CreateUser(*user)

	graph := &models.Graph{Id: 1, Edges: make([]*models.Edge, 0), Nodes: make([]*models.Node, 0)}
	user.Graphs = append(user.Graphs, graph)

	graph, err := testedDB.CreateGraph(*user, *graph)
	assert.Nil(t, err)
	assert.Equal(t, 1, int(graph.Id)) // we don't drop tables everytime, so id may be different

	graph, err = testedDB.CreateGraph(*user, *graph)
	assert.Nil(t, err)
	assert.Equal(t, 2, int(graph.Id))
}

func TestCreateNode(t *testing.T) {
	CleanUp()
	user := &models.User{}
	user, _ = testedDB.CreateUser(*user)

	graph := &models.Graph{}
	graph, _ = testedDB.CreateGraph(*user, *graph)

	node := &models.Node{Id: 1}

	node, err := testedDB.CreateNode(*graph, *node)
	assert.Nil(t, err)
	assert.Equal(t, 1, int(node.Id))

	node, err = testedDB.CreateNode(*graph, *node)
	assert.Nil(t, err)
	assert.Equal(t, 2, int(node.Id))
}

func TestCreateEdge(t *testing.T) {
	CleanUp()
	user := &models.User{}
	user, _ = testedDB.CreateUser(*user)

	graph := &models.Graph{}
	graph, _ = testedDB.CreateGraph(*user, *graph)

	edge := &models.Edge{Id: 1}

	edge, err := testedDB.CreateEdge(*graph, *edge)
	assert.Nil(t, err)
	assert.Equal(t, 1, int(edge.Id))

	edge, err = testedDB.CreateEdge(*graph, *edge)
	assert.Nil(t, err)
	assert.Equal(t, 2, int(edge.Id))
}

func TestUpdateUser(t *testing.T) {
	CleanUp()
	user := &models.User{
		Id:       1,
		Login:    "TestUpdate",
		Email:    "1",
		Password: "1",
		Graphs:   make([]*models.Graph, 0),
	}
	user, _ = testedDB.CreateUser(*user)

	graph := &models.Graph{Id: 239, Nodes: nil, Edges: nil}
	user.Graphs = append(user.Graphs, graph)

	userInDb, _ := testedDB.GetUserByLogin("TestUpdate")
	assert.Equal(t, 0, len(userInDb.Graphs))

	err := testedDB.UpdateUser(*user)
	assert.Nil(t, err)

	userInDb, err = testedDB.GetUserByLogin("TestUpdate")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(userInDb.Graphs))

	user.Login = "otherLogin"
	err = testedDB.UpdateUser(*user)
	assert.Nil(t, err)

	user.Id = -1
	err = testedDB.UpdateUser(*user)
	assert.NotNil(t, err)
}

func TestUpdateGraph(t *testing.T) {
	CleanUp()

	user := &models.User{
		Login: "TestUpdateGraph",
	}
	user, _ = testedDB.CreateUser(*user)
	graph := &models.Graph{
		Id:               1,
		Name:             "TestUpdateGraph1",
		Nodes:            make([]*models.Node, 0),
		Edges:            make([]*models.Edge, 0),
		DefaultNodeColor: models.ColorFromInt(0),
		DefaultEdgeColor: models.ColorFromInt(0),
	}
	graph, _ = testedDB.CreateGraph(*user, *graph)

	node := &models.Node{}
	node, err := testedDB.CreateNode(*graph, *node)
	assert.Nil(t, err)
	edge := &models.Edge{Description: "d", Color: models.ColorFromInt(0)}
	edge, err = testedDB.CreateEdge(*graph, *edge)
	assert.Nil(t, err)

	graph.Nodes = append(graph.Nodes, node)
	graph.Edges = append(graph.Edges, edge)

	graphInDB, err := testedDB.GetGraph(1)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(graphInDB.Nodes))

	err = testedDB.UpdateGraph(*graph)
	assert.Nil(t, err)

	graphInDB, err = testedDB.GetGraph(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(graphInDB.Nodes))
	assert.Equal(t, edge, graphInDB.Edges[0])
}

func TestMain(m *testing.M) {
	Setup()
	code := m.Run()
	TearDown()
	os.Exit(code)
}
