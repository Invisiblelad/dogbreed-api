package config

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "os"
    "time"
)

func ConnectDB() *mongo.Client {
    // Read MongoDB connection details from environment variables
    mongoURL := os.Getenv("MONGO_URL")
    mongoPort := os.Getenv("MONGO_PORT")
    mongoUser := os.Getenv("MONGO_USERNAME")
    mongoPass := os.Getenv("MONGO_PASSWORD")

    if mongoURL == "" || mongoPort == "" || mongoUser == "" || mongoPass == "" {
        log.Fatal("MONGO_URL, MONGO_PORT, MONGO_USERNAME, and MONGO_PASSWORD environment variables are required")
    }

    // Construct MongoDB connection URI
    mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoUser, mongoPass, mongoURL, mongoPort)

    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB With Authentication")
    return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    // Read database name from environment variable
    databaseName := os.Getenv("MONGO_DATABASE")
    if databaseName == "" {
        log.Fatal("MONGO_DATABASE environment variable is required")
    }

    collection := client.Database(databaseName).Collection(collectionName)
    return collection
}