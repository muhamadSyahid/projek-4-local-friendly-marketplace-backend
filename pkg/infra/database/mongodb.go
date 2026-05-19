package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

// InitMongoDB initializes MongoDB connection
func InitMongoDB() error {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		dbName = "pade_marketplace"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return err
	}

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return err
	}

	MongoClient = client
	MongoDB = client.Database(dbName)

	log.Println("MongoDB connected successfully")
	return nil
}

// DisconnectMongoDB closes the MongoDB connection
func DisconnectMongoDB() error {
	if MongoClient == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return MongoClient.Disconnect(ctx)
}

// GetDatabase returns the MongoDB database instance
func GetDatabase() *mongo.Database {
	return MongoDB
}
