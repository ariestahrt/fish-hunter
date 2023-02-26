package samples

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Feature struct {
	Text string `json:"text,omitempty" bson:"text,omitempty"`
	HTML string `json:"html,omitempty" bson:"html,omitempty"`
	// dont show this field in json
	CSS  string `json:"-" bson:"css,omitempty"`
}

type Domain struct {
	Id         	primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Ref_Dataset primitive.ObjectID 	`json:"ref_dataset" bson:"ref_dataset,omitempty"`
	URL			string 			   	`json:"url,omitempty" bson:"url,omitempty"`
	Brands	 	[]string           	`json:"brands,omitempty" bson:"brands,omitempty"`
	Language 	string 		   		`json:"language,omitempty" bson:"language,omitempty"`
	Details 	string 		   		`json:"details,omitempty" bson:"details,omitempty"`
	Type		string 		   		`json:"type,omitempty" bson:"type,omitempty"`
	Features	Feature           	`json:"features,omitempty" bson:"features,omitempty"`
	ScreenshotPath string 		   	`json:"screenshot_path,omitempty" bson:"screenshot_path,omitempty"`
	CreatedAt  	primitive.DateTime 	`json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  	primitive.DateTime 	`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt   primitive.DateTime 	`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type UseCase interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Update(id string, data *Domain) (Domain, error)
}

type Repository interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	CountTotal() (int64, error)
	GetTotalBetweenDates(startDate, endDate time.Time) (int64, error)
	Update(id string, data *Domain) (Domain, error)
}