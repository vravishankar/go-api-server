package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"../helper"
	"../model"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret")

// Auth - Authorization Function
func Auth(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var c model.Client
	err := decoder.Decode(&c)
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["key"] = c.APIKey
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString(mySigningKey)

	helper.RespondWithJSON(w, http.StatusOK, tokenString)
}
