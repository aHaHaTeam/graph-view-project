package api

//
//import (
//	"bytes"
//	"encoding/json"
//	"graph-view-project/models"
//	"io"
//	"log"
//	"net/http"
//	"strconv"
//)
//
//var address = "localhost:8080"
//var client *http.Client
//
//func LoadUser(userId int) *models.User {
//	response := makeRequest("GET", "/api/user/"+strconv.Itoa(userId), nil)
//	var user models.User
//	unmarshalBody(response, &user)
//	return &user
//}
//
//func LoadGraph(graphId int) models.Graph {
//	response := makeRequest("GET", "/api/user/"+strconv.Itoa(userId), nil)
//	var user models.User
//	unmarshalBody(response, &user)
//	return &user
//}
//
//func LoadNode(graphId int) models.Node {
//	response := makeRequest("GET", "/api/user/"+strconv.Itoa(userId), nil)
//	var user models.User
//	unmarshalBody(response, &user)
//	return &user
//}
//
//func LoadEdge(graphId int) models.Edge {
//	response := makeRequest("GET", "/api/user/"+strconv.Itoa(userId), nil)
//	var user models.User
//	unmarshalBody(response, &user)
//	return &user
//}
//
//func makeRequest(method, urlPath string, body interface{}) *http.Response {
//	requestBody, err := json.Marshal(body)
//	if err != nil {
//		log.Println(err)
//	}
//
//	request, err := http.NewRequest(method, address+urlPath, bytes.NewBuffer(requestBody))
//	if err != nil {
//		log.Println(err)
//	}
//
//	log.Println(address + urlPath)
//	response, err := client.Do(request)
//	if err != nil {
//		log.Println(err)
//		return nil
//	}
//	return response
//}
//
//func unmarshalBody(response *http.Response, result any) {
//	bodyBytes, err := io.ReadAll(response.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = json.Unmarshal(bodyBytes, &result)
//	if err != nil {
//		log.Println("[Unmarshal error]", err)
//	}
//}
