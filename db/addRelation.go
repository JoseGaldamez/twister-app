package db

import (
	"context"
	"time"
	"twister/app/models"
)

func AddRelation(relation models.Relation) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twisterdb")
	collection := db.Collection("relations")

	_, err := collection.InsertOne(ctx, relation)

	if err != nil {
		return false, err
	}

	return true, nil

}
