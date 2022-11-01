package response

import (
	"fish-hunter/businesses/jobs"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Ref_Url    primitive.ObjectID `json:"-"`
	HTTPStatus int                `json:"http_status,omitempty" bson:"http_status,omitempty"`
	SaveStatus string             `json:"save_status,omitempty" bson:"save_status,omitempty"`
	Details    string             `json:"details,omitempty" bson:"details,omitempty"`
	Worker     string             `json:"worker,omitempty" bson:"worker,omitempty"`
	CreatedAt  primitive.DateTime `json:"created_at,omitempty"`
	UpdatedAt  primitive.DateTime `json:"updated_at,omitempty"`
	DeletedAt   primitive.DateTime `json:"delete_at,omitempty"`
}

func FromDomain(domain jobs.Domain) Job {
	return Job{
		Id:         domain.Id,
		Ref_Url:    domain.Ref_Url,
		HTTPStatus: domain.HTTPStatus,
		SaveStatus: domain.SaveStatus,
		Details:    domain.Details,
		Worker:     domain.Worker,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}

func FromDomainArray(domain []jobs.Domain) []Job {
	var res []Job
	for _, value := range domain {
		res = append(res, FromDomain(value))
	}
	return res
}