package urls

import "go.mongodb.org/mongo-driver/bson/primitive"

type Domain struct {
    Id       		primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Ref_Source		primitive.ObjectID `json:"-"`
    Url 			string             `json:"url,omitempty" validate:"required"`
    Status  string			 `json:"status,omitempty"`
    Source_Url 		string             `json:"source_url,omitempty" validate:"required" bson:"source_url,omitempty"` 
    Source_Name 	string             `json:"source_name,omitempty" validate:"required" bson:"source_name,omitempty"` 
    CreatedAt		primitive.DateTime `json:"created_at,omitempty"`
    UpdatedAt		primitive.DateTime `json:"updated_at,omitempty"`
	DeleteAt		primitive.DateTime `json:"delete_at,omitempty"`
}

type UseCase interface {
	GetAll() ([]Domain, error)
	FetchUrl(source string) ([]Domain, error)
    GetByID(id string) (Domain, error)
}

type Repository interface {
	GetAll() ([]Domain, error)
    Save(domain Domain) (Domain, error)
    GetByID(id string) (Domain, error)
}