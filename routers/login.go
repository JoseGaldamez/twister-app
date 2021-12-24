package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"twister/app/db"
	"twister/app/jwt"
	"twister/app/models"
)

func Login(response http.ResponseWriter, request *http.Request) {

	response.Header().Add("content-type", "application/json")
	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		http.Error(response, "User/Password invalid"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(response, "Email is required", 400)
		return
	}

	document, exists := db.TryLogin(user.Email, user.Password)

	if !exists {
		http.Error(response, "User/Password invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(response, "Error generating JWT"+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)

	// This is the response that you see in Endpoint
	json.NewEncoder(response).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(response, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
