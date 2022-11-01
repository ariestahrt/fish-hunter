package requests

import (
	"errors"
	"fish-hunter/businesses/datasets"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatasetValidateRequest struct {
	Status string   `json:"status,omitempty" bson:"status,omitempty" form:"status,omitempty" validate:"required"`
	ScreenshotPath string   `json:"screenshot_path,omitempty" bson:"screenshot_path,omitempty" form:"screenshot_path,omitempty"`
	IsTweeted string   `json:"is_tweeted,omitempty" bson:"is_tweeted,omitempty" form:"is_tweeted,omitempty" validate:"required"`
}

func (req *DatasetValidateRequest) Validate() error {
	validate := validator.New()
	if validate.Struct(req) != nil {
		return errors.New("some of the fields are not valid")
	}
	return nil
}

func (req *DatasetValidateRequest) ToDomain(id string) *datasets.Domain {
	ObjId,_ := primitive.ObjectIDFromHex(id)
	return &datasets.Domain{
		Id:      ObjId,
		Status: req.Status,
		ScreenshotPath: req.ScreenshotPath,
	}
}