package routers

import (
	"io"
	"net/http"
	"os"
	"strings"
	"twister/app/db"
	"twister/app/models"
)

func UploadAvatar(repsonse http.ResponseWriter, request *http.Request) {
	file, handle, _ := request.FormFile("avatar")

	var extension = strings.Split(handle.Filename, ".")[1]

	var fileLocation string = "uploads/avatars/" + IDUser + "." + extension

	f, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(repsonse, "Error while saving image "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(repsonse, "Error while coping image "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension

	status, err = db.UpdateUser(user, IDUser)

	if err != nil || !status {
		http.Error(repsonse, "Error saving avatar in DB"+err.Error(), http.StatusBadRequest)
		return
	}

	repsonse.Header().Set("Content-Type", "application/json")
	repsonse.WriteHeader(http.StatusCreated)

}
