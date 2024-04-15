package test

import (
	"bytes"
	"encoding/json"
	"graph-view-project/database"
	"graph-view-project/models"
	"graph-view-project/server/handlers"
	"graph-view-project/server/routes"
	"graph-view-project/server/utils"
	"image/color"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var db database.DataBase = &database.MockDB{}

var address string
var server *httptest.Server
var router = mux.NewRouter()
var client *http.Client

func Setup() {
	db = &database.MockDB{}
	utils.JWTKey = []byte("testKey")
	_ = db.Connect("test")

	router = mux.NewRouter()
	routes.AddRoutes(router, &db)

	server = httptest.NewServer(router)
	client = server.Client()

	addr, err := url.Parse(server.URL)
	if err != nil {
		log.Println(err)
	}
	address = addr.String()
}

func TearDown() {
	_ = db.Disconnect()

	//server.Close()
}

func CleanUp() {
	TearDown()
	Setup()
}

func makeRequest(method, urlPath string, body interface{}, isAuthenticatedRequest bool, user AuthInput) *http.Response {
	requestBody, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
	}

	request, err := http.NewRequest(method, address+urlPath, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println(err)
	}

	if isAuthenticatedRequest {
		authRequest(request, user)
	}
	log.Println(address + urlPath)
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil
	}
	return response
}

func unmarshalBody(response *http.Response, result any) {
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		log.Println("[Unmarshal error]", err)
	}
}

type AuthInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func getToken(user AuthInput) []*http.Cookie {
	response := makeRequest("POST", "/login", user, false, user)
	return response.Cookies()
}

func authRequest(request *http.Request, user AuthInput) {
	cookies := getToken(user)
	request.AddCookie(cookies[0])
}

var auth = AuthInput{"user", "password"}

func TestApiAuth(t *testing.T) {
	CleanUp()

	response := makeRequest("GET", "/", nil, true, auth)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestApiUser(t *testing.T) {
	CleanUp()

	var expectedUser, actualUser, actualUser1 *models.User

	expectedUser = models.NewUser(
		0,
		"testLogin1",
		"testEmail1",
		"testPassword1",
		make([]*models.Graph, 0),
	)

	response := makeRequest("POST", "/api/user",
		handlers.CreateUserRequestBody{User: *expectedUser}, true, auth)
	unmarshalBody(response, &actualUser)

	expectedUser, err := db.GetUserByLogin(expectedUser.Login)
	if err != nil {
		log.Println("[Get user by login error]", err)
	}

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assertUsersEqual(t, *expectedUser, *actualUser)

	response = makeRequest("GET", "/api/user/"+strconv.Itoa(expectedUser.Id), nil, true, auth)
	unmarshalBody(response, &actualUser1)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assertUsersEqual(t, *expectedUser, *actualUser1)

	expectedUser.Graphs = append(expectedUser.Graphs, generateGraph(10, &db, *expectedUser))

	response = makeRequest("PUT", "/api/user/"+strconv.Itoa(expectedUser.Id),
		handlers.UpdateUserRequestBody{User: *expectedUser}, true, auth)
	unmarshalBody(response, &actualUser1)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assertUsersEqual(t, *expectedUser, *actualUser1)
}

func TestApiGraph(t *testing.T) {
	CleanUp()

	var expectedGraph, actualGraph, actualGraph1 *models.Graph

	var user models.User
	response := makeRequest("GET", "/api/user/"+strconv.Itoa(1), nil, true, auth)
	unmarshalBody(response, &user)
	expectedGraph = generateGraph(10, &db, user)

	response = makeRequest("POST", "/api/graph",
		handlers.CreateGraphRequestBody{User: user, Graph: *expectedGraph}, true, auth)
	unmarshalBody(response, &actualGraph)

	expectedGraph, err := db.GetGraph(actualGraph.Id)
	if err != nil {
		log.Println("[Get user by login error]", err)
	}

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assertGraphsEqual(t, *expectedGraph, *actualGraph)

	response = makeRequest("GET", "/api/graph/"+strconv.Itoa(expectedGraph.Id), nil, true, auth)
	unmarshalBody(response, &actualGraph1)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assertGraphsEqual(t, *expectedGraph, *actualGraph1)

	expectedGraph.Nodes = append(expectedGraph.Nodes, generateNode(&db, *expectedGraph))

	response = makeRequest("PUT", "/api/graph/"+strconv.Itoa(expectedGraph.Id),
		handlers.UpdateGraphRequestBody{Graph: *expectedGraph}, true, auth)
	unmarshalBody(response, &actualGraph1)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assertGraphsEqual(t, *expectedGraph, *actualGraph1)
}

func TestApiNode(t *testing.T) {
	CleanUp()

	var actualNode *models.Node

	var user models.User

	response := makeRequest("GET", "/api/user/"+strconv.Itoa(1), nil, true, auth)
	unmarshalBody(response, &user)
	graph := generateGraph(0, &db, user)

	node1 := models.NewNode(0, make([]int, 0),
		"testNodeName2", []byte("testNodeData2"),
		1.0, color.RGBA{0, 0, 0, 0}, models.Square)

	node2 := models.NewNode(0, make([]int, 0),
		"testNodeName2", []byte("testNodeData2"),
		1.0, color.RGBA{0, 0, 0, 0}, models.Circle)

	// POST node1
	response = makeRequest("POST", "/api/node",
		handlers.CreateNodeRequestBody{Graph: *graph, Node: *node1}, true, auth)
	unmarshalBody(response, &actualNode)
	node1.Id = actualNode.Id

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assertNodesEqual(t, *node1, *actualNode)

	// POST node2
	response = makeRequest("POST", "/api/node",
		handlers.CreateNodeRequestBody{Graph: *graph, Node: *node2}, true, auth)
	unmarshalBody(response, &actualNode)
	node2.Id = actualNode.Id

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assertNodesEqual(t, *node2, *actualNode)

	// GET node1
	response = makeRequest("GET", "/api/node/"+strconv.Itoa(node1.Id), nil, true, auth)
	unmarshalBody(response, &actualNode)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assertNodesEqual(t, *node1, *actualNode)

	node1.Edges = append(node1.Edges, node2.Id)
	node2.Edges = append(node1.Edges, node1.Id)

	// PUT node1
	response = makeRequest("PUT", "/api/node/"+strconv.Itoa(node1.Id),
		handlers.UpdateNodeRequestBody{Node: *node1}, true, auth)
	unmarshalBody(response, &actualNode)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assertNodesEqual(t, *node1, *actualNode)

	// PUT node2
	response = makeRequest("PUT", "/api/node/"+strconv.Itoa(node2.Id),
		handlers.UpdateNodeRequestBody{Node: *node2}, true, auth)
	unmarshalBody(response, &actualNode)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assertNodesEqual(t, *node2, *actualNode)
}

func TestApiEdge(t *testing.T) {
	CleanUp()

	var actualNode *models.Node

	var user models.User

	response := makeRequest("GET", "/api/user/"+strconv.Itoa(1), nil, true, auth)
	unmarshalBody(response, &user)
	graph := generateGraph(0, &db, user)

	node1 := models.NewNode(0, make([]int, 0),
		"testNodeName2", []byte("testNodeData2"),
		1.0, color.RGBA{0, 0, 0, 0}, models.Square)

	node2 := models.NewNode(0, make([]int, 0),
		"testNodeName2", []byte("testNodeData2"),
		1.0, color.RGBA{0, 0, 0, 0}, models.Circle)

	// POST node1
	response = makeRequest("POST", "/api/node",
		handlers.CreateNodeRequestBody{Graph: *graph, Node: *node1}, true, auth)
	unmarshalBody(response, &actualNode)
	node1.Id = actualNode.Id

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assertNodesEqual(t, *node1, *actualNode)

	// POST node2
	response = makeRequest("POST", "/api/node",
		handlers.CreateNodeRequestBody{Graph: *graph, Node: *node2}, true, auth)
	unmarshalBody(response, &actualNode)
	node2.Id = actualNode.Id

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assertNodesEqual(t, *node2, *actualNode)

	edge := models.NewEdge(0, node1.Id, node2.Id,
		"testEdgeName1", "testNodeDescription1",
		0.1, color.RGBA{0, 0, 0, 0}, models.Arrow)
	var actualEdge *models.Edge

	// POST edge
	response = makeRequest("POST", "/api/edge",
		handlers.CreateEdgeRequestBody{Graph: *graph, Edge: *edge}, true, auth)
	unmarshalBody(response, &actualEdge)
	edge.Id = actualEdge.Id

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assertEdgesEqual(t, *edge, *actualEdge)

	// GET edge
	response = makeRequest("GET", "/api/edge/"+strconv.Itoa(edge.Id), nil, true, auth)
	unmarshalBody(response, &actualEdge)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assertEdgesEqual(t, *edge, *actualEdge)

	// PUT edge
	response = makeRequest("PUT", "/api/edge/"+strconv.Itoa(edge.Id),
		handlers.UpdateEdgeRequestBody{Edge: *edge}, true, auth)
	unmarshalBody(response, &actualEdge)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assertEdgesEqual(t, *edge, *actualEdge)
}
