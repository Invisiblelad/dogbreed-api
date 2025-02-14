package handlers

import (
    "encoding/json"
	
	"github.com/invisiblelad/DogBreedApi/models"
	"github.com/invisiblelad/DogBreedApi/repositiories"
    "go.mongodb.org/mongo-driver/bson"
    "net/http"
    "github.com/go-chi/chi/v5"
)

type DogBreedHandler struct {
    Repo *repositories.DogBreedRepository
}

func NewDogBreedHandler(repo *repositories.DogBreedRepository) *DogBreedHandler {
    return &DogBreedHandler{Repo: repo}
}

func (h *DogBreedHandler) CreateDogBreed(w http.ResponseWriter, r *http.Request) {
    var dogBreed models.DogBreed
    json.NewDecoder(r.Body).Decode(&dogBreed)
    createdDogBreed, err := h.Repo.Create(&dogBreed)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(createdDogBreed)
}

func (h *DogBreedHandler) GetAllDogBreeds(w http.ResponseWriter, r *http.Request) {
    dogBreeds, err := h.Repo.Getall()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(dogBreeds)
}

func (h *DogBreedHandler) GetDogBreedByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    dogBreed, err := h.Repo.FindByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(dogBreed)
}

func (h *DogBreedHandler) UpdateDogBreed(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var dogBreed models.DogBreed
    json.NewDecoder(r.Body).Decode(&dogBreed)
    err := h.Repo.Update(id, &dogBreed)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (h *DogBreedHandler) DeleteDogBreed(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    err := h.Repo.Delete(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}


func (h *DogBreedHandler)DeleteManyDogBreed(w http.ResponseWriter, r *http.Request){
    var filter bson.M 

    json.NewDecoder(r.Body).Decode(&filter)

    err := h.Repo.DeleteMany(filter)

    if err !=nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}