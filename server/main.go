package main

import (
	"graph-view-project/database"
	"graph-view-project/server/routes"
	"graph-view-project/server/utils"
	"log"
	"mime"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	utils.JWTKey = []byte(os.Getenv("JWT_KEY"))

	err = mime.AddExtensionType(".css", "text/css; charset=utf-8")
	if err != nil {
		return
	}
	err = mime.AddExtensionType(".js", "application/javascript; charset=utf-8")
	if err != nil {
		return
	}
	var db database.DataBase
	db = &database.MockDB{}
	_ = db.Connect("postgres")

	router := mux.NewRouter()
	routes.AddRoutes(router, &db)

	log.Println("Listening on port :8080")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
