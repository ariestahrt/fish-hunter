package users

import (
	"fish-hunter/businesses/users"
	"time"
)

type User struct {
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Roles     []string `json:"roles"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt time.Time   `json:"deleted_at"`
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		Username:  domain.Username,
		Email:     domain.Email,
		Password:  domain.Password,
		Roles:     domain.Roles,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (u *User) ToDomain() users.Domain {
	return users.Domain{
		Username:  u.Username,
		Email:     u.Email,
		Roles:     u.Roles,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}