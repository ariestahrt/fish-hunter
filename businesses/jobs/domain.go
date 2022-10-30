package jobs

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	DeleteAt   primitive.DateTime `json:"delete_at,omitempty"`
}

type UseCase interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
}

type Repository interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
}