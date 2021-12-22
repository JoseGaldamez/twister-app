package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name,omitempty"`
	Lastname string             `bson:"lastName" json:"lastName,omitempty"`
	BirthDay time.Time          `bson:"birthDay" json:"birthDay,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
	Avatar   string             `bson:"avatar" json:"avatar,omitempty"`
	Banner   string             `bson:"banner" json:"banner,omitempty"`
	Biografy string             `bson:"biografy" json:"biografy,omitempty"`
	Location string             `bson:"location" json:"location,omitempty"`
	WebSite  string             `bson:"webSite" json:"webSite,omitempty"`
}
