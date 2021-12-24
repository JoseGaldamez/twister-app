package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(ID string, UserID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twisterdb")
	collection := db.Collection("tweets")

	ObjID, _ := primitive.ObjectIDFromHex(ID)

	query := bson.M{
		"_id":    ObjID,
		"userid": UserID,
	}

	_, err := collection.DeleteOne(ctx, query)
	return err

}
