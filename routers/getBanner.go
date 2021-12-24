package routers

import (
	"io"
	"net/http"
	"os"
	"twister/app/db"
)

func GetBanner(response http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(response, "ID is required", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(response, "Profile not found", http.StatusBadRequest)
		return
	}

	opefile, err := os.Open("uploads/banners/" + profile.Banner)

	if err != nil {
		http.Error(response, "Image no found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(response, opefile)

	if err != nil {
		http.Error(response, "Error while coping image", http.StatusBadRequest)
	}

}
