package controllers

import (
	"encoding/json"
	"graph-view-project/database"
	"graph-view-project/models"
	"graph-view-project/server/utils"
	"log"
	"net/http"
)

func Signup(db *database.DataBase) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var user models.User

		if err := decoder.Decode(&user); err != nil {
			w.Header().Add("success", "Invalid username or password")
			log.Println(err)
			return
		}

		var err error
		user.Password, err = utils.GenerateHashPassword(user.Password)

		if err != nil {
			w.Header().Add("success", "Invalid password")
			log.Println(err)
			return
		}

		err = (*db).CreateUser(user)

		if err != nil {
			w.Header().Add("success", "User with this username already exists")
			log.Println(err)
			return
		}

		w.Header().Add("success", "User created")
	}
}
