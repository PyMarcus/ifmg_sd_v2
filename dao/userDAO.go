package dao

import (
	"context"
	"log"

	"github.com/PyMarcus/distribuidos/models"
	"github.com/PyMarcus/distribuidos/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	users []*models.User
}

func NewUser() *User {
	return &User{}
}

func (u User) GetAllUsers() []*models.User {
	conn := utils.ConnectionFactory()
	defer conn.Disconnect(context.TODO())

	database := conn.Database("user")
	collection := database.Collection("users")

	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Panic(err)
		return nil
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var result bson.M

		if err := cursor.Decode(&result); err != nil {
			log.Println(err)
			return nil
		}

		users := models.NewUser(result["_id"].(primitive.ObjectID),
			result["name"].(string),
			result["email"].(string),
		)
		u.users = append(u.users, users)
	}

	return u.users
}

func (u User) CreateUser(user *models.User) {
	conn := utils.ConnectionFactory()
	defer conn.Disconnect(context.TODO())

	database := conn.Database("user")
	collection := database.Collection("users")

	_, err := collection.InsertOne(context.TODO(), *user)

	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("Created!")
	return
}

func (u User) UpdateUser(oldName, oldEmail, name, email string) {
	log.Println(email, oldEmail, oldName, name)
	conn := utils.ConnectionFactory()
	defer conn.Disconnect(context.TODO())

	database := conn.Database("user")
	collection := database.Collection("users")

	filter := bson.M{"name": oldEmail, "email": oldName}

	update := bson.M{"$set": bson.M{"email": email, "name": name}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("Updated!")
	return
}

func (u User) DeleteUser(name, email string){
	conn := utils.ConnectionFactory()
	defer conn.Disconnect(context.TODO())

	database := conn.Database("user")
	collection := database.Collection("users")

	filter := bson.M{"name": email, "email": name}


	_, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("Removed!")
	return
}