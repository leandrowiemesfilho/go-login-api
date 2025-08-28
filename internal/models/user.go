package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `json:"user_id" bson:"_id"`
	FirstName    string             `json:"first_name" bson:"first_name"`
	LastName     string             `json:"last_name" bson:"last_name"`
	Email        string             `json:"email" bson:"email"`
	PhoneNumber  string             `json:"phone_number" bson:"phone_number"`
	Password     string             `json:"password" bson:"password"`
	Roles        []string           `json:"roles" bson:"roles"`
	CreationDate time.Time          `json:"creation_date" bson:"creation_date"`
	UpdateDate   time.Time          `json:"update_date" bson:"update_date"`
}
