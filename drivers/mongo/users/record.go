package users

import (
	"fish-hunter/businesses/users"
	"time"
)

type User struct {
	ID		string    `bson:"_id,omitempty" json:"_id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	IsActive  bool     `json:"is_active" bson:"is_active"`
	Name	  string   `json:"name"`
	Phone	  string   `json:"phone"`
	University string  `json:"university"`
	Position   string  `json:"position"`
	Proposal   string  `json:"proposal"`
	Roles     []string `json:"roles"`
	CreatedAt time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time   `json:"deleted_at" bson:"deleted_at"`
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID: 	   domain.ID,
		Username:  domain.Username,
		Email:     domain.Email,
		Password:  domain.Password,
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
		res = append(res, *FromDomain(&value))
	}
	return res
}

func (u *User) ToDomain() users.Domain {
	return users.Domain{
		ID: 	   u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		IsActive:  u.IsActive,
		Name:	   u.Name,
		Phone:	   u.Phone,
		University: u.University,
		Position:   u.Position,
		Proposal:   u.Proposal,
		Roles:     u.Roles,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}

func ToDomainArray(u *[]User) []users.Domain {
	var res []users.Domain
	for _, value := range *u {
		res = append(res, value.ToDomain())
	}
	return res
}