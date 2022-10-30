package jobs

import (
	"fish-hunter/businesses/jobs"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Ref_Url    primitive.ObjectID `json:"ref_url" bson:"ref_url"`
	Url		   string             `json:"url" bson:"url"`
	HTTPStatus int    `json:"http_status" bson:"http_status"`
	SaveStatus string `json:"save_status" bson:"save_status"`
	Details    string `json:"details" bson:"details"`
	Worker     string `json:"worker" bson:"worker"`
	CreatedAt  primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt  primitive.DateTime `json:"updated_at" bson:"updated_at"`
	DeleteAt   primitive.DateTime `json:"delete_at" bson:"delete_at"`
}

func FromDomain(domain jobs.Domain) Job {
	return Job{
		Id:         domain.Id,
		Ref_Url:    domain.Ref_Url,
		Url: 	  	domain.Url,
		HTTPStatus: domain.HTTPStatus,
		SaveStatus: domain.SaveStatus,
		Details:    domain.Details,
		Worker:     domain.Worker,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeleteAt:   domain.DeleteAt,
	}
}

func FromDomainArray(domain []jobs.Domain) []Job {
	var res []Job
	for _, value := range domain {
		res = append(res, FromDomain(value))
	}
	return res
}

func (job *Job) ToDomain() jobs.Domain {
	return jobs.Domain{
		Id:         job.Id,
		Ref_Url:    job.Ref_Url,
		Url: 		job.Url,
		HTTPStatus: job.HTTPStatus,
		SaveStatus: job.SaveStatus,
		Details:    job.Details,
		Worker:     job.Worker,
		CreatedAt:  job.CreatedAt,
		UpdatedAt:  job.UpdatedAt,
		DeleteAt:   job.DeleteAt,
	}
}

func ToDomainArray(job *[]Job) []jobs.Domain {
	var res []jobs.Domain
	for _, value := range *job {
		res = append(res, value.ToDomain())
	}
	return res
}