package routes

import (
	"graph-view-project/database"
	"graph-view-project/server/handlers"
	"graph-view-project/server/middleware"

	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router, db *database.DataBase) {
	addApiRoutes(router, db)

	router.HandleFunc("/login", handlers.Login(db)).Methods("POST")
	router.HandleFunc("/signup", handlers.Signup(db)).Methods("POST")

	// Pages
	router.HandleFunc("/login", handlers.LoginGet).Methods("GET")
	router.HandleFunc("/logout", handlers.Logout).Methods("GET")

	// Private pages
	router.HandleFunc("/", middleware.AuthUser(handlers.Home)).Methods("GET")

	// Resources
	router.PathPrefix("/").HandlerFunc(handlers.ServeStatic).Methods("GET")
}

func addApiRoutes(router *mux.Router, db *database.DataBase) {
	// API
	router.HandleFunc("/api/user/{id:[0-9]+}", middleware.AuthUser(handlers.GetUser(db))).Methods("GET")
	router.HandleFunc("/api/user/{id:[0-9]+}", middleware.AuthUser(handlers.GetUser(db))).Methods("GET")
	router.HandleFunc("/api/graph/{id:[0-9]+}", middleware.AuthUser(handlers.GetGraph(db))).Methods("GET")
	router.HandleFunc("/api/edge/{id:[0-9]+}", middleware.AuthUser(handlers.GetEdge(db))).Methods("GET")
	router.HandleFunc("/api/node/{id:[0-9]+}", middleware.AuthUser(handlers.GetNode(db))).Methods("GET")

	router.HandleFunc("/api/user/{id:[0-9]+}", middleware.AuthUser(handlers.UpdateUser(db))).Methods("PUT")
	router.HandleFunc("/api/graph/{id:[0-9]+}", middleware.AuthUser(handlers.UpdateGraph(db))).Methods("PUT")
	router.HandleFunc("/api/edge/{id:[0-9]+}", middleware.AuthUser(handlers.UpdateEdge(db))).Methods("PUT")
	router.HandleFunc("/api/node/{id:[0-9]+}", middleware.AuthUser(handlers.UpdateNode(db))).Methods("PUT")

	router.HandleFunc("/api/user", middleware.AuthUser(handlers.CreateUser(db))).Methods("POST")
	router.HandleFunc("/api/graph", middleware.AuthUser(handlers.CreateGraph(db))).Methods("POST")
	router.HandleFunc("/api/edge", middleware.AuthUser(handlers.CreateEdge(db))).Methods("POST")
	router.HandleFunc("/api/node", middleware.AuthUser(handlers.CreateNode(db))).Methods("POST")
}
