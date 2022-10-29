package users

import "time"

type Domain struct {
	ID        string    `bson:"_id,omitempty" json:"_id"`
	Username  string
	Email     string
	Password  string
	NewPassword  string
	IsActive  bool
	Name	  string
	Phone	  string
	University string
	Position   string // Student, Lecturer, Staff
	Proposal   string
	Roles     []string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UseCase interface {
	Register(domain *Domain) (Domain, error)
	Login(domain *Domain) (string, error)
	UpdateProfile(domain *Domain) (Domain, error)
	UpdatePassword(domain *Domain) (Domain, error)
	GetProfile(id string) (Domain, error)
	GetAllUsers() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Update(domain *Domain) (Domain, error)
	Delete(id string) (Domain, error)
	Logout(token string) error
}

type Repository interface {
	Register(domain *Domain) (Domain, error)
	Login(domain *Domain) (Domain, error)
	UpdateProfile(domain *Domain) (Domain, error)
	UpdatePassword(domain *Domain) (Domain, error)
	GetProfile(id string) (Domain, error)
	GetAllUsers() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Update(domain *Domain) (Domain, error)
	Delete(id string) (Domain, error)
}