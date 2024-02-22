package routes

import (
	"graph-view-project/server/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/signup", controllers.Signup).Methods("POST")

	router.HandleFunc("/login", LoginGet).Methods("GET")
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/logout", controllers.Logout).Methods("GET")

	router.PathPrefix("/").HandlerFunc(ServeStatic).Methods("GET")
}

func LoginGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/static/login.html")
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/static/"+r.URL.Path)
}
