package main

import (
    "github.com/go-chi/chi/v5"
	"github.com/invisiblelad/DogBreedApi/config"
	"github.com/invisiblelad/DogBreedApi/repositiories"
	"github.com/invisiblelad/DogBreedApi/handlers"
    "github.com/go-chi/chi/v5/middleware"
    "net/http"
    "os"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    client := config.ConnectDB()

    // Read collection name from environment variable
    collectionName := os.Getenv("MONGO_COLLECTION")
    if collectionName == "" {
        panic("MONGO_COLLECTION environment variable is required")
    }

    dogBreedCollection := config.GetCollection(client, collectionName)
    dogBreedRepo := repositories.NewDogBreedRepository(dogBreedCollection)
    dogBreedHandler := handlers.NewDogBreedHandler(dogBreedRepo)

    r.Post("/dogbreeds", dogBreedHandler.CreateDogBreed)
    r.Get("/dogbreeds", dogBreedHandler.GetAllDogBreeds)
    r.Get("/dogbreeds/{id}", dogBreedHandler.GetDogBreedByID)
    r.Put("/dogbreeds/{id}", dogBreedHandler.UpdateDogBreed)
    r.Delete("/dogbreeds/{id}", dogBreedHandler.DeleteDogBreed)
    r.Delete("/dogbreeds/many", dogBreedHandler.DeleteManyDogBreed)

    http.ListenAndServe(":8080", r)
}