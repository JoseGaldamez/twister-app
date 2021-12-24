package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"twister/app/db"
)

func ReadAllUsers(response http.ResponseWriter, request *http.Request) {

	typeuser := request.URL.Query().Get("type")
	page := request.URL.Query().Get("page")
	search := request.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(response, "Page is required", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := db.ReadAllUsers(IDUser, pag, search, typeuser)

	if !status {
		http.Error(response, "Something went wrong while reading users", http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)

	json.NewEncoder(response).Encode(result)

}
