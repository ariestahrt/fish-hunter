package sources

import "go.mongodb.org/mongo-driver/bson/primitive"

type Domain struct {
    Id		primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Name	string             `json:"name,omitempty" validate:"required"`
	Url		string             `json:"url,omitempty" validate:"required"`
}

type Repository interface {
	GetByName(name string) (Domain, error)
}