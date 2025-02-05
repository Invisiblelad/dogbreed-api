package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DogBreed struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Name        string             `json:"name" bson:"name"`
    Description string             `json:"description" bson:"description"`
    Origin      string             `json:"origin" bson:"origin"`
}