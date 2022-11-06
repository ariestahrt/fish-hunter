package jobs_test

import (
	"fish-hunter/businesses/jobs"
	"fish-hunter/businesses/jobs/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	jobRepository mocks.Repository
	jobService    jobs.UseCase
	jobDomain     jobs.Domain
)

/*
type Domain struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Ref_Url    primitive.ObjectID `json:"-" bson:"ref_url,omitempty"`
	Url		   string             `json:"url,omitempty" bson:"url,omitempty"`
	HTTPStatus int                `json:"http_status,omitempty" bson:"http_status,omitempty"`
	SaveStatus string             `json:"save_status,omitempty" bson:"save_status,omitempty"`
	Details    string             `json:"details,omitempty" bson:"details,omitempty"`
	Worker     string             `json:"worker,omitempty" bson:"worker,omitempty"`
	CreatedAt  primitive.DateTime `json:"created_at,omitempty"`
	UpdatedAt  primitive.DateTime `json:"updated_at,omitempty"`
	DeletedAt   primitive.DateTime `json:"delete_at,omitempty"`
}
*/

func TestMain(m *testing.M) {
	jobService = jobs.NewJobUseCase(&jobRepository)
	jobDomain = jobs.Domain{
		Id:         primitive.NewObjectID(),
		Ref_Url:    primitive.NewObjectID(),
		Url:        "https://www.google.com",
		HTTPStatus: 200,
		SaveStatus: "success",
		Details:    "success",
		Worker:     "worker",
		CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:  primitive.NewDateTimeFromTime(time.Now()),
		DeletedAt:  primitive.NewDateTimeFromTime(time.Time{}),
	}
	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Test Case 1 | Get All", func(t *testing.T) {
		jobRepository.On("GetAll").Return([]jobs.Domain{jobDomain}, nil).Once()
		result, err := jobService.GetAll()

		assert.Equal(t, []jobs.Domain{jobDomain}, result)
		assert.Nil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Test Case 1 | Get By ID", func(t *testing.T) {
		jobRepository.On("GetByID", jobDomain.Id.Hex()).Return(jobDomain, nil).Once()
		result, err := jobService.GetByID(jobDomain.Id.Hex())

		assert.Equal(t, jobDomain, result)
		assert.Nil(t, err)
	})
}