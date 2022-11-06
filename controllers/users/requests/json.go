package requests

import (
	"errors"
	"fish-hunter/businesses/users"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRegister struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=12"`
	Name	 string `json:"name" validate:"required"`
	Phone	 string `json:"phone" validate:"required"`
	University string `json:"university" validate:"required"`
	Position   string `json:"position" validate:"required"`
	Proposal   string `json:"proposal" validate:"required"`
	Roles   []string `json:"roles" validate:"required"`
}

func (u *UserRegister) ToDomain() *users.Domain {
	return &users.Domain{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Name:	 u.Name,
		Phone:	 u.Phone,
		University: u.University,
		Position:   u.Position,
		Proposal:   u.Proposal,
		Roles:    u.Roles,
	}
}

func (u *UserRegister) Validate() error {
	validate := validator.New()
	if validate.Struct(u) != nil {
		return errors.New("some of the fields are not valid")
	}
	return nil
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

// For User Update
type UserUpdateProfile struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Name	 string `json:"name" validate:"required"`
	Phone	 string `json:"phone" validate:"required"`
	University string `json:"university" validate:"required"`
	Position   string `json:"position" validate:"required"`
	Proposal   string `json:"proposal" validate:"required"`
}

func (u *UserUpdateProfile) ToDomain(id string) *users.Domain {
	ObjId, _ := primitive.ObjectIDFromHex(id)
	return &users.Domain{
		Id : ObjId,
		Username: u.Username,
		Email:    u.Email,
		Name:	 u.Name,
		Phone:	 u.Phone,
		University: u.University,
		Position:   u.Position,
		Proposal:   u.Proposal,
	}
}

func (u *UserUpdateProfile) Validate() error {
	validate := validator.New()
	if validate.Struct(u) != nil {
		return errors.New("some of the fields are not valid")
	}
	return nil
}

// For Update Password

type UserUpdatePassword struct {
	OldPassword string `json:"old_password" validate:"required,min=8,max=12"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=12"`
}

func (u *UserUpdatePassword) ToDomain(id string) *users.Domain {
	ObjId, _ := primitive.ObjectIDFromHex(id)
	return &users.Domain{
		Id : ObjId,
		Password: u.OldPassword,
		NewPassword: u.NewPassword,
	}
}

func (u *UserUpdatePassword) Validate() error {
	validate := validator.New()
	if validate.Struct(u) != nil {
		return errors.New("some of the fields are not valid")
	}
	return nil
}

// For Update By Admin

type UserUpdateByAdmin struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	Name	 string `json:"name" validate:"required"`
	Phone	 string `json:"phone" validate:"required"`
	IsActive bool `json:"is_active" validate:"required"`
	University string `json:"university" validate:"required"`
	Position   string `json:"position" validate:"required"`
	Proposal   string `json:"proposal" validate:"required"`
	Roles   []string `json:"roles" validate:"required"`
}

func (u *UserUpdateByAdmin) ToDomain(id string) *users.Domain {
	ObjId, _ := primitive.ObjectIDFromHex(id)
	return &users.Domain{
		Id : ObjId,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Name:	 u.Name,
		Phone:	 u.Phone,
		IsActive: u.IsActive,
		University: u.University,
		Position:   u.Position,
		Proposal:   u.Proposal,
		Roles:    u.Roles,
	}
}

func (u *UserUpdateByAdmin) Validate() error {
	validate := validator.New()
	if validate.Struct(u) != nil {
		return errors.New("some of the fields are not valid")
	}
	return nil
}