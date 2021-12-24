package db

import (
	"context"
	"fmt"
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
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})

	query = append(query, bson.M{
		"$unwind": "$tweet",
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

	cursor, err := collection.Aggregate(ctx, query)

	if err != nil {
		fmt.Println("Aggregate ejecutado con errores")
	}

	var result []models.ReadTweetsFollowers

	err = cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}

	cursor.Close(ctx)

	return result, true

}
