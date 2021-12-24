package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"twister/app/db"
)

func ReadTweetsFollowers(response http.ResponseWriter, request *http.Request) {

	if len(request.URL.Query().Get("page")) < 1 {
		http.Error(response, "Page is required", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil {
		http.Error(response, "Page must be a number", http.StatusBadRequest)
		return
	}

	result, ok := db.ReadTweetsFollowers(IDUser, page)

	if !ok {
		http.Error(response, "Error while reading tweets", http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type", " application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(result)

}
