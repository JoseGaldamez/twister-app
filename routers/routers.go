package routers

import (
	"encoding/json"
	"net/http"

	"twister/app/db"
	"twister/app/models"
)

// Register is a function to create a user in DB
func Register(response http.ResponseWriter, request *http.Request) {

	db.CheckConnection()

	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		http.Error(response, "Error on data sent"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(response, "Email is required", 400)
		return
	}

	if len(user.Password) == 6 {
		http.Error(response, "Password must be at least 6 characters", 400)
		return
	}

	_, found, _ := db.CheckUserExists(user.Email)

	if found {
		http.Error(response, "Email already exists", 400)
		return
	}

	_, status, err := db.CreateRegister(user)

	if err != nil {
		http.Error(response, "Something went wrong while writting user"+err.Error(), 400)
		return
	}

	if status {
		http.Error(response, "Users is not active"+err.Error(), 400)
		return
	}

	response.WriteHeader(http.StatusCreated)

}
