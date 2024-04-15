package handlers

import (
	"net/http"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/static/login.html")
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/static/index.html")
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/static"+r.URL.Path)
}
