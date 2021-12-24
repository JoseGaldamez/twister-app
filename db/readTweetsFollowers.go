package db

import (
	"context"
	"time"
	"twister/app/models"

	"go.mongodb.org/mongo-driver/bson"
)

func ReadTweetsFollowers(ID string, page int) ([]models.ReadTweetsFollowers, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twisterdb")
	collection := db.Collection("relations")

	skip := ((page - 1) * 20)

	query := make([]bson.M, 0)

	query = append(query, bson.M{"$match": bson.M{"userid": ID}})
	query = append(query, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "userrealationid",
			"foreingField": "userid",
			"as":           "usertweets",
		},
	})

	query = append(query, bson.M{
		"$unwind": "$usertweets",
	})

	query = append(query, bson.M{
		"$sort": bson.M{
			"date": -1,
		},
	})

	query = append(query, bson.M{
		"$skip": skip,
	})

	query = append(query, bson.M{
		"$limit": 20,
	})

	cursor, _ := collection.Aggregate(ctx, query)

	var results []models.ReadTweetsFollowers

	err := cursor.All(ctx, &results)
	if err != nil {
		return results, false
	}

	return results, true

}
