package db

import (
	"context"
	"time"
	"twister/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Update a user with an ID
func UpdateUser(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twisterdb")
	collection := db.Collection("users")

	updateRegister := make(map[string]interface{})

	if len(user.Name) > 0 {
		updateRegister["name"] = user.Name
	}
	if len(user.Lastname) > 0 {
		updateRegister["lastname"] = user.Lastname
	}

	updateRegister["birthDay"] = user.BirthDay

	if len(user.Avatar) > 0 {
		updateRegister["avatar"] = user.Avatar
	}

	if len(user.Biografy) > 0 {
		updateRegister["biografy"] = user.Biografy
	}

	if len(user.Location) > 0 {
		updateRegister["location"] = user.Location
	}
	if len(user.Banner) > 0 {
		updateRegister["banner"] = user.Banner
	}

	if len(user.WebSite) > 0 {
		updateRegister["webSite"] = user.WebSite
	}

	// updateRegister containts new objet with fields updated
	updateQuery := bson.M{
		"$set": updateRegister,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	// point to user owner of ID
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := collection.UpdateOne(ctx, filter, updateQuery)

	if err != nil {
		return false, err
	}

	return true, nil

}
