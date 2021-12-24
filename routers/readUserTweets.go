package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"twister/app/db"
)

func ReadUserTweets(response http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(response, "Id is required", http.StatusBadRequest)
		return
	}

	if len(request.URL.Query().Get("page")) < 1 {
		http.Error(response, "page parameter is required", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(request.URL.Query().Get("page"))

	if err != nil {
		http.Error(response, "page must be a number > 0", http.StatusBadRequest)
		return
	}

	pag := int64(page)

	results, ok := db.ReadUserTweets(ID, pag)

	if !ok {
		http.Error(response, "Error reading tweets", http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(results)

}
