package routes

import (
	"graph-view-project/database"
	"graph-view-project/server/controllers"

	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router, db *database.DataBase) {
	router.HandleFunc("/login", controllers.Login(db)).Methods("POST")
	router.HandleFunc("/signup", controllers.Signup(db)).Methods("POST")

	router.HandleFunc("/login", controllers.LoginGet).Methods("GET")
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/logout", controllers.Logout).Methods("GET")

	router.PathPrefix("/").HandlerFunc(controllers.ServeStatic).Methods("GET")
}
