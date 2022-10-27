package response

import (
	"fish-hunter/businesses/users"
	"time"
)

type User struct {
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt time.Time   `json:"deleted_at"`
}

func FromDomain(domain users.Domain) User {
	return User{
		Username:  domain.Username,
		Email:     domain.Email,
		Roles:     domain.Roles,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}