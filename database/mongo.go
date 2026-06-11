package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var StudentCollection *mongo.Collection

func ConnectDB() {

	uri := "mongodb+srv://go:go@cluster0.vvgdhov.mongodb.net/?appName=Cluster0"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(uri),
	)

	if err != nil {
		log.Fatal(err)
	}

	StudentCollection = client.
		Database("college").
		Collection("students")

	log.Println("✅ MongoDB Connected")
}