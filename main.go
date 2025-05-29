package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"GoWithMongoDB/controllers"
)

func main() {
	r := httprouter.New()

	// Get MongoDB client
	client, err := getClient()
	if err != nil {
		fmt.Printf("Failed to connect to MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(context.Background())

	// Initialize controller
	uc := controllers.NewUserController(client)
	r.GET("/user/:id", uc.GetUser)
	r.GET("/user", uc.GetUsers)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	fmt.Println("Server starting on http://localhost:9000")
	if err := http.ListenAndServe("localhost:9000", r); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}

// connection with mongodb
func getClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017") // Update to 27107 if needed
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!")
	return client, nil
}