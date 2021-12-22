package db

import (
	"context"
	"time"
	"twister/app/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateRegister(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twisterdb")
	collection := db.Collection("users")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
