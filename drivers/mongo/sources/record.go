package sources

import (
	"fish-hunter/businesses/sources"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Source struct {
	Id   primitive.ObjectID `json:"id"`
	Name string             `json:"name" validate:"required"`
	Url  string             `json:"url" validate:"required"`
}

func FromDomain(domain *sources.Domain) *Source {
	return &Source{
		Id:   domain.Id,
		Name: domain.Name,
		Url:  domain.Url,
	}
}

func (s *Source) ToDomain() sources.Domain {
	return sources.Domain{
		Id:   s.Id,
		Name: s.Name,
		Url:  s.Url,
	}
}