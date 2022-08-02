package databases

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx = context.TODO()

	// UserCollection exported MongoDB Collection
	UserCollection *mongo.Collection
	// BugTypeCollection exported MongoDB Collection
	BugTypeCollection *mongo.Collection
	// TicketCollection exported MongoDB Collection
	TicketCollection *mongo.Collection
	// VerificationCollection exported MongoDB Collection
	VerificationCollection *mongo.Collection

	//quickCollection *mongo.Collection
	//mongoClient *mongo.Client
)

// Init func
func Init() {
	ClientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err := mongo.Connect(ctx, ClientOptions)
	if err != nil {
		log.Println(err)
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Connected to MongoDB!")

	UserCollection = mongoClient.Database("bugTracker").Collection("users")
	BugTypeCollection = mongoClient.Database("bugTracker").Collection("bugtype")
	TicketCollection = mongoClient.Database("bugTracker").Collection("tickets")
	VerificationCollection = mongoClient.Database("bugTracker").Collection("verification")
}
