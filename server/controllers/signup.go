package controllers

import (
	"encoding/json"
	"graph-view-project/server/database"
	"graph-view-project/server/models"
	"graph-view-project/server/utils"
	"log"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User

	if err := decoder.Decode(&user); err != nil {
		w.Header().Add("success", "Invalid username or password")
		//http.Error(w, err.Error(), 400)
		log.Println(err)
		return
	}

	var err error
	user.Password, err = utils.GenerateHashPassword(user.Password)

	if err != nil {
		w.Header().Add("success", "Invalid password")
		//http.Error(w, "could not generate password hash", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = database.DB.CreateUser(user.Login, user.Email, user.Password)

	if err != nil {
		w.Header().Add("success", "User with this username already exists")
		//http.Error(w, "user already exists", http.StatusBadRequest)
		log.Println(err)
		return
	}

	w.Header().Add("success", "User created")
}
