package routers

import (
	"encoding/json"
	"net/http"
	"twister/app/db"
	"twister/app/models"
)

func CheckRelation(response http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")

	var relation models.Relation
	relation.UserID = IDUser
	relation.UserRelationID = ID

	var res models.ResponseSearchRelation

	status, err := db.CheckRelation(relation)

	if err != nil || !status {
		res.Status = false
	} else {
		res.Status = true
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)

	json.NewEncoder(response).Encode(res)

}
