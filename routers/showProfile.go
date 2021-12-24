package routers

import (
	"encoding/json"
	"net/http"
	"twister/app/db"
)

// ShowProfile return the user requested
func ShowProfile(response http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(response, "ID is required", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(response, "User not found"+err.Error(), 400)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)

	json.NewEncoder(response).Encode(profile)

}
