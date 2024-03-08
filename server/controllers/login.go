package controllers

import (
	"encoding/json"
	"graph-view-project/database"
	"graph-view-project/server/models"
	"graph-view-project/server/utils"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User

	if err := decoder.Decode(&user); err != nil {
		w.Header().Add("success", "Invalid login or password")
		//http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	var existingUser models.User

	existingUser, err := database.DB.GetUserByLogin(user.Login)

	if err != nil {
		w.Header().Add("success", "Invalid login or password")
		//http.Error(w, "user does not exist", http.StatusBadRequest)
		log.Println(err)
		return
	}

	err = utils.CompareHashPassword(user.Password, existingUser.Password)

	if err != nil {
		w.Header().Add("success", "Invalid login or password")
		//http.Error(w, "invalid password", http.StatusBadRequest)
		log.Println(err)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		UserId: 3, //existingUser.Id,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Login,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(utils.JWTKey)

	if err != nil {
		w.Header().Add("success", "Invalid login or password")
		//http.Error(w, "could not generate token", http.StatusInternalServerError)
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

	w.Header().Add("success", "User logged in")
}
