package handlers

import (
	"net/http"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/static/login.html")
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/static/"+r.URL.Path)
}
