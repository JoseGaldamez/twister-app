package db

import (
	"context"
	"log"
	"time"
	"twister/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Return tweet of an user
func ReadUserTweets(ID string, page int64) ([]*models.ReturnTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twisterdb")
	collection := db.Collection("tweets")

	var results []*models.ReturnTweet

	query := bson.M{
		"userid": ID,
	}

	option := options.Find()
	option.SetLimit(20)
	option.SetSort(bson.D{{Key: "date", Value: -1}})
	option.SetSkip((page - 1) * 20)

	cursor, err := collection.Find(ctx, query, option)

	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var tweet models.ReturnTweet
		err := cursor.Decode(&tweet)

		if err != nil {
			return results, false
		}

		results = append(results, &tweet)
	}

	return results, true
}
