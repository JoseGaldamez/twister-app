package db

import (
	"context"
	"fmt"
	"time"
	"twister/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twisterdb")
	collection := db.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	query := bson.M{
		"_id": objID,
	}

	err := collection.FindOne(ctx, query).Decode(&profile)
	profile.Password = ""

	if err != nil {
		fmt.Println("User not found " + err.Error())
		return profile, err
	}

	return profile, nil

}
