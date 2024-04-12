package routes

import (
	"graph-view-project/database"
	"graph-view-project/server/handlers"
	"graph-view-project/server/middleware"

	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router, db *database.DataBase) {
	router.HandleFunc("/login", handlers.Login(db)).Methods("POST")
	router.HandleFunc("/signup", handlers.Signup(db)).Methods("POST")

	router.HandleFunc("/login", handlers.LoginGet).Methods("GET")
	router.HandleFunc("/", middleware.AuthUser(handlers.Home)).Methods("GET")
	router.HandleFunc("/logout", handlers.Logout).Methods("GET")

	router.PathPrefix("/").HandlerFunc(middleware.AuthUser(handlers.ServeStatic)).Methods("GET")

	router.HandleFunc("api/getUser", middleware.AuthUser(handlers.GetUser(db))).Methods("GET")
	router.HandleFunc("api/getGraph", middleware.AuthUser(handlers.GetGraph(db))).Methods("GET")
	router.HandleFunc("api/getEdge", middleware.AuthUser(handlers.GetEdge(db))).Methods("GET")
	router.HandleFunc("api/getNode", middleware.AuthUser(handlers.GetNode(db))).Methods("GET")

	router.HandleFunc("api/updateUser", middleware.AuthUser(handlers.UpdateUser(db))).Methods("POST")
	router.HandleFunc("api/updateGraph", middleware.AuthUser(handlers.UpdateGraph(db))).Methods("POST")
	router.HandleFunc("api/updateEdge", middleware.AuthUser(handlers.UpdateEdge(db))).Methods("POST")
	router.HandleFunc("api/updateNode", middleware.AuthUser(handlers.UpdateNode(db))).Methods("POST")

	router.HandleFunc("api/createUser", middleware.AuthUser(handlers.CreateUser(db))).Methods("POST")
	router.HandleFunc("api/createGraph", middleware.AuthUser(handlers.CreateGraph(db))).Methods("POST")
	router.HandleFunc("api/createEdge", middleware.AuthUser(handlers.CreateEdge(db))).Methods("POST")
	router.HandleFunc("api/createNode", middleware.AuthUser(handlers.CreateNode(db))).Methods("POST")
}
