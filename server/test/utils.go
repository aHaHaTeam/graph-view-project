package test

import (
	"graph-view-project/database"
	"graph-view-project/models"
	"image/color"
	"log"
	"math/rand"
	"strconv"

	"github.com/stretchr/testify/assert"
)

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func generateNode(db *database.DataBase, graph models.Graph) *models.Node {
	node := models.NewNode(
		0,
		make([]int, 0),
		"testNodeName"+strconv.Itoa(rand.Int()),
		[]byte("testNodeData"+strconv.Itoa(rand.Int())),
		1.0,
		color.RGBA{
			R: uint8(randInt(0, 256)),
			G: uint8(randInt(0, 256)),
			B: uint8(randInt(0, 256)),
			A: uint8(randInt(0, 256)),
		},
		models.Square,
	)
	node, err := (*db).CreateNode(graph, *node)
	if err != nil {
		log.Println("[Node generating error]", err)
	}
	return node
}

func generateEdge(begin, end *models.Node, db *database.DataBase, graph models.Graph) *models.Edge {
	edge := models.NewEdge(
		0,
		begin.Id,
		end.Id,
		"testEdgeName"+strconv.Itoa(rand.Int()),
		"testEdgeDescription"+strconv.Itoa(rand.Int()),
		0.1,
		color.RGBA{
			R: uint8(randInt(0, 256)),
			G: uint8(randInt(0, 256)),
			B: uint8(randInt(0, 256)),
			A: uint8(randInt(0, 256)),
		},
		models.Arrow,
	)
	edge, err := (*db).CreateEdge(graph, *edge)
	if err != nil {
		log.Println("[Edge generating error]", err)
	}
	return edge
}

func generateGraph(n int, db *database.DataBase, user models.User) *models.Graph {
	graph := models.NewGraph(
		0,
		"testGraphName"+strconv.Itoa(rand.Int()),
		"testGraphDescription"+strconv.Itoa(rand.Int()),
		make([]*models.Node, 0),
		make([]*models.Edge, 0),
		1.0,
		color.RGBA{R: 123, G: 123, B: 123, A: 123},
		models.Square,
		0.1,
		color.RGBA{R: 123, G: 123, B: 123, A: 123},
		models.Arrow,
	)

	graph, err := (*db).CreateGraph(user, *graph)
	if err != nil {
		log.Println("[Graph generating error]", err)
	}

	graph.Nodes = append(graph.Nodes, generateNode(db, *graph))
	for i := 1; i < n; i++ {
		graph.Nodes = append(graph.Nodes, generateNode(db, *graph))
		graph.Edges = append(graph.Edges, generateEdge(graph.Nodes[randInt(0, i)], graph.Nodes[i], db, *graph))
	}
	err = (*db).UpdateGraph(*graph)
	if err != nil {
		log.Println("[Graph generating error]", err)
	}

	return graph
}

func assertUsersEqual(t assert.TestingT, expected models.User, actual models.User) {
	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.Login, actual.Login)
	assert.Equal(t, expected.Email, actual.Email)
	assert.Equal(t, expected.Password, actual.Password)
	assert.Equal(t, len(expected.Graphs), len(actual.Graphs))
	for i, g := range expected.Graphs {
		assertGraphsEqual(t, *g, *actual.Graphs[i])
	}
}

func assertGraphsEqual(t assert.TestingT, expected models.Graph, actual models.Graph) {
	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Description, actual.Description)
	for i, n := range expected.Nodes {
		assert.Equal(t, n.Id, actual.Nodes[i].Id)
	}
	for i, e := range expected.Edges {
		assert.Equal(t, e.Id, actual.Edges[i].Id)
	}
}

func assertNodesEqual(t assert.TestingT, expected models.Node, actual models.Node) {
	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Data, actual.Data)
	for i, n := range expected.Edges {
		assert.Equal(t, n, actual.Edges[i])
	}
	assert.Equal(t, expected.Color, actual.Color)
	assert.Equal(t, expected.Shape, actual.Shape)
	assert.Equal(t, expected.Size, actual.Size)
}

func assertEdgesEqual(t assert.TestingT, expected models.Edge, actual models.Edge) {
	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Begin, actual.Begin)
	assert.Equal(t, expected.End, actual.End)
	assert.Equal(t, expected.Color, actual.Color)
	assert.Equal(t, expected.Shape, actual.Shape)
	assert.Equal(t, expected.Width, actual.Width)
}
