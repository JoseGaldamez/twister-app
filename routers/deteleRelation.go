package routers

import (
	"net/http"
	"twister/app/db"
	"twister/app/models"
)

func DeleteRelation(response http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(response, "Id is required", http.StatusBadRequest)
		return
	}

	var relation models.Relation

	relation.UserID = IDUser
	relation.UserRelationID = ID

	status, err := db.DeleteRelation(relation)

	if err != nil {
		http.Error(response, "Error while deleting"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(response, "Relation no deleted", http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusCreated)

}
