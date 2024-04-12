package middleware

import (
	"graph-view-project/server/handlers"
	"graph-view-project/server/utils"
	"log"
	"net/http"
	"strconv"
)

func AuthUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")

		if err != nil {
			handlers.LoginGet(w, r)
			log.Println(err)
			return
		}

		claims, err := utils.ParseToken(cookie.Value)

		if err != nil {
			handlers.LoginGet(w, r)
			log.Println(err)
			return
		}

		r.Header.Add("success", "home page")
		r.Header.Add("email", strconv.Itoa(claims.UserId))

		next(w, r)
	}
}
