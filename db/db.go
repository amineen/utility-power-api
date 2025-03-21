package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Define ANSI escape codes for colors
	orange := "\033[33m"
	reset := "\033[0m"

	fmt.Printf("%sConnected to MongoDB!%s\n", orange, reset)
	Client = client
}

func GetDatabase() *mongo.Database {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Client.Database(os.Getenv("MONGODB_DATABASE"))
}

func GetCollection(name string) *mongo.Collection {
	return GetDatabase().Collection(name)
}

func GetCustomersCollection() *mongo.Collection {
	return GetCollection("customers")
}

func GetVendorSalesSummaryCollection() *mongo.Collection {
	return GetCollection("vendor-sales-summaries")
}
