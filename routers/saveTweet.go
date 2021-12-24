package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"twister/app/db"
	"twister/app/models"

	"go.mongodb.org/mongo-driver/bson"
)

func SaveTweet(response http.ResponseWriter, request *http.Request) {

	var message models.Tweet

	err := json.NewDecoder(request.Body).Decode(&message)

	if err != nil {
		http.Error(response, "Imposible read body request", 400)
		return
	}

	regiterTweet := models.AddTweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.AddTweet(regiterTweet)

	if err != nil {
		http.Error(response, "Something went wrong while saving "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(response, "Couldn't save tweet ", 400)
		return
	}

	msg := bson.M{
		"msg": "Tweet saved",
		"ok":  true,
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(msg)

}
