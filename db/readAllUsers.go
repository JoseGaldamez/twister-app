package db

import (
	"context"
	"fmt"
	"time"
	"twister/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadAllUsers(ID string, page int64, search string, usertype string) ([]*models.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twisterdb")
	collection := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, added bool

	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)

		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var relation models.Relation
		relation.UserID = ID
		relation.UserRelationID = user.ID.Hex()

		added = false
		found, _ = CheckRelation(relation)

		if usertype == "new" && !found {
			added = true
		}

		if usertype == "follow" && found {
			added = true
		}

		if relation.UserRelationID == ID {
			added = false
		}

		if added {
			user.Password = ""
			user.Biografy = ""
			user.WebSite = ""
			user.Location = ""
			user.Email = ""
			user.Banner = ""

			results = append(results, &user)

		}

	}

	err = cursor.Err()

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cursor.Close(ctx)

	return results, true

}
