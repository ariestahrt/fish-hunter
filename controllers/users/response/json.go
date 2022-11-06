package response

import (
	"fish-hunter/businesses/users"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id		primitive.ObjectID    `bson:"_id,omitempty" json:"_id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	IsActive  bool     `json:"is_active"`
	Name	  string   `json:"name"`
	Phone	  string   `json:"phone"`
	University string  `json:"university"`
	Position   string  `json:"position"`
	Proposal   string  `json:"proposal"`
	Roles     []string `json:"roles"`
	CreatedAt primitive.DateTime   `json:"created_at"`
	UpdatedAt primitive.DateTime   `json:"updated_at"`
	DeletedAt primitive.DateTime   `json:"deleted_at"`
}

func FromDomain(domain users.Domain) User {
	return User{
		Id: 	   domain.Id,
		Username:  domain.Username,
		Email:     domain.Email,
		IsActive:  domain.IsActive,
		Name:	   domain.Name,
		Phone:	   domain.Phone,
		University: domain.University,
		Position:   domain.Position,
		Proposal:   domain.Proposal,
		Roles:     domain.Roles,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func FromDomainArray(domain []users.Domain) []User {
	var res []User
	for _, value := range domain {
		res = append(res, FromDomain(value))
	}
	return res
}