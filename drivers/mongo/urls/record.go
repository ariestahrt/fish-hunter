package urls

import (
	"fish-hunter/businesses/urls"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Url struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Ref_Source  primitive.ObjectID `json:"-"`
	Url         string             `json:"url,omitempty" validate:"required"`
	Executed    bool               `json:"executed,omitempty"`
	Source_Url  string             `json:"source_url,omitempty" validate:"required" bson:"source_url,omitempty"`
	Source_Name string             `json:"source_name,omitempty" validate:"required" bson:"source_name,omitempty"`
	CreatedAt   primitive.DateTime `json:"created_at,omitempty", bson:"created_at,omitempty"`
	UpdatedAt   primitive.DateTime `json:"updated_t,omitempty", bson:"updated_at,omitempty"`
	DeleteAt    primitive.DateTime `json:"delete_at,omitempty, bson:"delete_at,omitempty"`
}

func FromDomain(domain urls.Domain) Url {
	return Url{
		Id:          domain.Id,
		Ref_Source:  domain.Ref_Source,
		Url:         domain.Url,
		Executed:    domain.Executed,
		Source_Url:  domain.Source_Url,
		Source_Name: domain.Source_Name,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeleteAt:    domain.DeleteAt,
	}
}

func FromDomainArray(domain []urls.Domain) []Url {
	var res []Url
	for _, value := range domain {
		res = append(res, FromDomain(value))
	}
	return res
}

func (url *Url) ToDomain() urls.Domain {
	return urls.Domain{
		Id:          url.Id,
		Ref_Source:  url.Ref_Source,
		Url:         url.Url,
		Executed:    url.Executed,
		Source_Url:  url.Source_Url,
		Source_Name: url.Source_Name,
		CreatedAt:   url.CreatedAt,
		UpdatedAt:   url.UpdatedAt,
		DeleteAt:    url.DeleteAt,
	}
}

func ToDomainArray(url *[]Url) []urls.Domain {
	var res []urls.Domain
	for _, value := range *url {
		res = append(res, value.ToDomain())
	}
	return res
}