package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id    primitive.ObjectID `json:"_id" bson:"_id"`
	Email string `json:"email" bson:"email"`
	Name  string `json:"name" bson:"name"`
}

func NewUser(id primitive.ObjectID, email, name string) *User {
	return &User{
		Id: id,
		Email: email,
		Name:  name,
	}
}