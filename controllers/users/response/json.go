package response

import (
	"fish-hunter/businesses/users"
	"time"
)

type User struct {
	ID		string    `bson:"_id,omitempty" json:"_id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	IsActive  bool     `json:"is_active"`
	Name	  string   `json:"name"`
	Phone	  string   `json:"phone"`
	University string  `json:"university"`
	Position   string  `json:"position"`
	Proposal   string  `json:"proposal"`
	Roles     []string `json:"roles"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt time.Time   `json:"deleted_at"`
}

func FromDomain(domain users.Domain) User {
	return User{
		ID: 	   domain.ID,
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