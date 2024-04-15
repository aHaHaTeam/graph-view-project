package handlers

import (
	"encoding/json"
	"graph-view-project/database"
	"graph-view-project/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CreateUserRequestBody struct {
	User models.User `json:"user"`
}

func CreateUser(db *database.DataBase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var body CreateUserRequestBody
		err := dec.Decode(&body)

		if err != nil {
			log.Printf("Couldn't unmarshal a user interface. error: %s\n", err)
			respondWithUnmarshallError(w, err)
			return
		}

		newUser, err := (*db).CreateUser(body.User)
		if err != nil {
			log.Printf("Couldn't insert a user into database. error: %s\n", err)
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusCreated, newUser)
	}
}

type CreateGraphRequestBody struct {
	User  models.User  `json:"user"`
	Graph models.Graph `json:"graph"`
}

func CreateGraph(db *database.DataBase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var body CreateGraphRequestBody
		err := dec.Decode(&body)

		if err != nil {
			respondWithUnmarshallError(w, err)
			return
		}

		newGraph, err := (*db).CreateGraph(body.User, body.Graph)
		if err != nil {
			log.Printf("Couldn't insert a graph into database. error: %s\n", err)
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusCreated, newGraph)
	}
}

type CreateNodeRequestBody struct {
	Graph models.Graph `json:"graph"`
	Node  models.Node  `json:"node"`
}

func CreateNode(db *database.DataBase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var body CreateNodeRequestBody
		err := dec.Decode(&body)

		if err != nil {
			respondWithUnmarshallError(w, err)
			return
		}

		newNode, err := (*db).CreateNode(body.Graph, body.Node)
		if err != nil {
			log.Printf("Couldn't insert a node into database. error: %s\n", err)
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusCreated, newNode)
	}
}

type CreateEdgeRequestBody struct {
	Graph models.Graph `json:"graph"`
	Edge  models.Edge  `json:"edge"`
}

func CreateEdge(db *database.DataBase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var body CreateEdgeRequestBody
		err := dec.Decode(&body)

		if err != nil {
			respondWithUnmarshallError(w, err)
			return
		}

		newEdge, err := (*db).CreateEdge(body.Graph, body.Edge)
		if err != nil {
			log.Printf("Couldn't insert a edge into database. error: %s\n", err)
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusCreated, newEdge)
	}
}

func GetUser(db *database.DataBase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GetUser")
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid Graph ID")
			return
		}

		user, err := (*db).GetUser(id)
		if err != nil {
			log.Printf("User with id: \"%s\" requested but not found.\n", vars["id"])
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, user)
	}
}

func GetGraph(db *database.DataBase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid Graph ID")
			return
		}

		graph, err := (*db).GetGraph(id)
		if err != nil {
			log.Printf("Graph with id: \"%s\" requested but not found.\n", vars["id"])
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, graph)
	}
}

func GetEdge(db *database.DataBase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid Edge ID")
			return
		}

		edge, err := (*db).GetEdge(id)
		if err != nil {
			log.Printf("Edge with id: \"%s\" requested but not found.\n", vars["id"])
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, edge)
	}
}

func GetNode(db *database.DataBase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid Node ID")
			return
		}

		node, err := (*db).GetNode(id)
		if err != nil {
			log.Printf("Node with id: \"%s\" requested but not found.\n", vars["id"])
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, node)
	}
}

type UpdateUserRequestBody struct {
	User models.User `json:"user"`
}

func UpdateUser(db *database.DataBase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var body UpdateUserRequestBody
		err := dec.Decode(&body)

		if err != nil {
			respondWithUnmarshallError(w, err)
			return
		}

		if err = (*db).UpdateUser(body.User); err != nil {
			log.Printf("Couldn't update a user with login: \"%s\". error: %s\n", body.User.Login, err)
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, body.User)
	}
}

type UpdateGraphRequestBody struct {
	Graph models.Graph `json:"graph"`
}

func UpdateGraph(db *database.DataBase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var body UpdateGraphRequestBody
		err := dec.Decode(&body)

		if err != nil {
			respondWithUnmarshallError(w, err)
			return
		}

		if err = (*db).UpdateGraph(body.Graph); err != nil {
			log.Printf("Couldn't update a graph with id: %d. error: %s\n", body.Graph.Id, err)
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, body.Graph)
	}
}

type UpdateEdgeRequestBody struct {
	Edge models.Edge `json:"edge"`
}

func UpdateEdge(db *database.DataBase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var body UpdateEdgeRequestBody
		err := dec.Decode(&body)

		if err != nil {
			respondWithUnmarshallError(w, err)
			return
		}

		if err = (*db).UpdateEdge(body.Edge); err != nil {
			log.Printf("Couldn't update a edge with id: %d. error: %s\n", body.Edge.Id, err)
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, body.Edge)
	}
}

type UpdateNodeRequestBody struct {
	Node models.Node `json:"node"`
}

func UpdateNode(db *database.DataBase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var body UpdateNodeRequestBody
		err := dec.Decode(&body)

		if err != nil {
			respondWithUnmarshallError(w, err)
			return
		}

		if err = (*db).UpdateNode(body.Node); err != nil {
			log.Printf("Couldn't update a node with id: %d. error: %s\n", body.Node.Id, err)
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, body.Node)
	}
}
