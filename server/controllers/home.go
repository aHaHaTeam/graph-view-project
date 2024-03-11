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
		LoginGet(w, r)
		log.Println(err)
		return
	}

	claims, err := utils.ParseToken(cookie.Value)

	if err != nil {
		LoginGet(w, r)
		log.Println(err)
		return
	}

	r.Header.Add("success", "home page")
	r.Header.Add("email", strconv.Itoa(claims.UserId))

	http.ServeFile(w, r, "./client/static/index.html")
}
