package handlers

import (
	"encoding/json"
	"graph-view-project/database"
	"graph-view-project/models"
	"graph-view-project/server/utils"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func Login(db *database.DataBase) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var user models.User

		if err := decoder.Decode(&user); err != nil {
			w.Header().Add("success", "Invalid login or password")
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("Decode error", err)
			return
		}
		var existingUser *models.User

		existingUser, err := (*db).GetUserByLogin(user.Login)

		if err != nil {
			w.Header().Add("success", "Invalid login or password")
			w.WriteHeader(http.StatusUnauthorized)
			log.Println(err)
			return
		}

		err = utils.CompareHashPassword(user.Password, existingUser.Password)

		if err != nil {
			w.Header().Add("success", "Invalid login or password")
			w.WriteHeader(http.StatusUnauthorized)
			hash, _ := utils.GenerateHashPassword(user.Password)
			log.Println(err, user.Login, user.Password, hash, existingUser.Password)
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &models.Claims{
			UserId: existingUser.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(utils.JWTKey)

		if err != nil {
			w.Header().Add("success", "Invalid login or password")
			w.WriteHeader(http.StatusUnauthorized)
			log.Println(err)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:       "token",
			Value:      tokenString,
			Path:       "/",
			Domain:     "localhost",
			Expires:    expirationTime,
			RawExpires: expirationTime.String(),
			MaxAge:     int(time.Until(expirationTime).Seconds()),
			Secure:     false,
			HttpOnly:   true,
		})

		log.Println("User logged in")
		w.Header().Add("success", "User logged in")
		w.WriteHeader(http.StatusOK)
	}
}
