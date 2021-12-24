package db

import (
	"context"
	"time"

	"twister/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTweet(tweet models.AddTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twisterdb")
	collection := db.Collection("tweets")

	register := bson.M{
		"userid":  tweet.UserID,
		"message": tweet.Message,
		"date":    tweet.Date,
	}

	result, err := collection.InsertOne(ctx, register)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.Hex(), true, nil

}
