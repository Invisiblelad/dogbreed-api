package repositories

import (
	"context"

	"github.com/invisiblelad/DogBreedApi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DogBreedRepository struct {
    Collection *mongo.Collection
}

func NewDogBreedRepository(collection *mongo.Collection) *DogBreedRepository {
	return &DogBreedRepository{Collection: collection}
}

func (r *DogBreedRepository)Create(dogBreed *models.DogBreed)(*models.DogBreed, error){
	dogBreed.ID = primitive.NewObjectID()

	_,err := r.Collection.InsertOne(context.Background(), dogBreed)

	if err !=nil{
		panic(err)
	}
	return dogBreed, err
}

func (r *DogBreedRepository)Getall()([]models.DogBreed,error){
	var dogBreeds []models.DogBreed

	cursor , err := r.Collection.Find(context.Background(),bson.M{})

	if err!=nil{
		panic(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()){
		var dogBreed models.DogBreed
		cursor.Decode(&dogBreed)

		dogBreeds = append(dogBreeds, dogBreed)
	}
	return dogBreeds,nil
}

func (r *DogBreedRepository)FindByID(id string)(*models.DogBreed, error){
	objectID , err  := primitive.ObjectIDFromHex(id)

	if err !=nil{
		panic(err)
	}
	var dogBreed models.DogBreed
	err = r.Collection.FindOne(context.Background(),bson.M{"_id": objectID}).Decode(&dogBreed)
	return &dogBreed, err
	
}

func (r *DogBreedRepository) Update(id string, dogBreed *models.DogBreed) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = r.Collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, bson.M{"$set": dogBreed})
    return err
}

func (r *DogBreedRepository) Delete(id string) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = r.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
    return err
}