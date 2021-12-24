package routers

import (
	"net/http"
	"twister/app/db"
)

func DeleteTweet(response http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(response, "Id is required", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, IDUser)

	if err != nil {
		http.Error(response, "Something went wrong while deleting", http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)

}
