package requests

import (
	"fish-hunter/businesses/datasets"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatasetValidateRequest struct {
	Status string   `json:"status,omitempty" bson:"status,omitempty"`
}

func (req *DatasetValidateRequest) ToDomain(id string) *datasets.Domain {
	ObjId,_ := primitive.ObjectIDFromHex(id)
	return &datasets.Domain{
		Id:      ObjId,
		Status: req.Status,
	}
}