package db

import (
	"context"
	"time"
	"twister/app/models"

	"go.mongodb.org/mongo-driver/bson"
)

func DeleteRelation(relation models.Relation) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twisterdb")
	collection := db.Collection("relations")

	query := bson.M{
		"userid":         relation.UserID,
		"userrelationid": relation.UserRelationID,
	}

	_, err := collection.DeleteOne(ctx, query)

	if err != nil {
		return false, err
	}

	return true, nil

}
