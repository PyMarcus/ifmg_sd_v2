package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectionFactory creates a connection with a NoSql mongo db database.
func ConnectionFactory() *mongo.Client {
	databaseClient := options.Client().ApplyURI(URL)

	client, err := mongo.Connect(context.TODO(), databaseClient)

	if err != nil {
		log.Println("FAIL TO CONNECT INTO DATABASE")
		log.Fatalln(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("FAIL TO CHECK CONNECTION")
		log.Fatalln(err)
	}

	log.Println("Connected to database")

	return client
}