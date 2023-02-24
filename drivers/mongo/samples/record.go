package samples

import (
	"fish-hunter/businesses/samples"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sample struct {
	Id         	primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Ref_Dataset primitive.ObjectID 	`json:"-" bson:"ref_dataset,omitempty"`
	URL			string 			   	`json:"url,omitempty" bson:"url,omitempty"`
	Brands	 	[]string           	`json:"brands,omitempty" bson:"brands,omitempty"`
	Language 	string 		   		`json:"language,omitempty" bson:"language,omitempty"`
	Details 	string 		   		`json:"details,omitempty" bson:"details,omitempty"`
	Type		string 		   		`json:"type,omitempty" bson:"type,omitempty"`
	Features	samples.Feature		`json:"features,omitempty" bson:"features,omitempty"`
	ScreenshotPath string 		   	`json:"screenshot_path,omitempty" bson:"screenshot_path,omitempty"`
	CreatedAt  	primitive.DateTime 	`json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  	primitive.DateTime 	`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt   primitive.DateTime 	`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

func FromDomain(domain *samples.Domain) Sample {
	return Sample{
		Id:         domain.Id,
		Ref_Dataset: domain.Ref_Dataset,
		URL: 		domain.URL,
		Brands: 	domain.Brands,
		Language: 	domain.Language,
		Details:    domain.Details,
		Type: 		domain.Type,
		Features: 	domain.Features,
		ScreenshotPath: domain.ScreenshotPath,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}

func FromDomainArray(domain []samples.Domain) []Sample {
	var res []Sample
	for _, value := range domain {
		res = append(res, FromDomain(&value))
	}
	return res
}

func (sample *Sample) ToDomain() samples.Domain {
	return samples.Domain{
		Id:         sample.Id,
		Ref_Dataset: sample.Ref_Dataset,
		URL: 		sample.URL,
		Brands: 	sample.Brands,
		Language: 	sample.Language,
		Details:    sample.Details,
		Type: 		sample.Type,
		Features: 	sample.Features,
		ScreenshotPath: sample.ScreenshotPath,
		CreatedAt:  sample.CreatedAt,
		UpdatedAt:  sample.UpdatedAt,
		DeletedAt:   sample.DeletedAt,
	}
}

func ToDomainArray(sample *[]Sample) []samples.Domain {
	var res []samples.Domain
	for _, value := range *sample {
		res = append(res, value.ToDomain())
	}
	return res
}