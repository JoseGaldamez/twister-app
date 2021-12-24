package routers

import (
	"encoding/json"
	"net/http"

	"twister/app/db"
	"twister/app/models"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateProfile(response http.ResponseWriter, request *http.Request) {
	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(response, "Incorrect data"+err.Error(), 400)
		return
	}
	var status bool
	status, err = db.UpdateUser(user, IDUser)

	if err != nil {
		http.Error(response, "Something went wrong while updating "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(response, "User not found", 400)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)

	msg := bson.M{"msg": "Profile updated"}

	json.NewEncoder(response).Encode(msg)

}
