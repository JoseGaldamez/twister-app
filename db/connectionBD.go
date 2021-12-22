package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN contain the connection to DB
var MongoCN = conectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://JoseGaldamez:Magodeoz1991@bdtwister.0emur.mongodb.net/twisterdb?retryWrites=true&w=majority")

func conectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connection succesfully")
	return client

}

// CheckConnection let me test connection with DB
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
