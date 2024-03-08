package controllers

import (
	"graph-view-project/server/utils"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		//http.ServeFile(w, r, "../client/static/login.html")
		log.Println(err)
		return
	}

	claims, err := utils.ParseToken(cookie.Value)

	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		//http.ServeFile(w, r, "../client/static/login.html")
		log.Println(err)
		return
	}

	r.Header.Add("success", "home page")
	r.Header.Add("email", strconv.Itoa(claims.UserId))

	http.ServeFile(w, r, "../client/static/index.html")
}
