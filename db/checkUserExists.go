package db

import (
	"context"
	"time"
	"twister/app/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := MongoCN.Database("twisterdb")
	collection := db.Collection("users")

	query := bson.M{"email": email}

	var result models.User
	err := collection.FindOne(ctx, query).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID

}
