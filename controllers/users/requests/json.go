package requests

import (
	"fish-hunter/businesses/users"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=12"`
	Roles   []string `json:"roles" validate:"required"`
}

func (u *User) ToDomain() *users.Domain {
	return &users.Domain{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Roles:    u.Roles,
	}
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=12"`
}

func (u *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u *UserLogin) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}