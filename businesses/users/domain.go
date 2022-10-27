package users

import "time"

type Domain struct {
	Username  string
	Email     string
	Password  string
	Roles     []string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UseCase interface {
	Register(domain *Domain) Domain
	Login(domain *Domain) (Domain, error)
	// GetAll() ([]Domain, error)
	// GetByID(id int) (Domain, error)
	// Update(domain Domain) (Domain, error)
	// Delete(domain Domain) (Domain, error)
}

type Repository interface {
	Register(domain *Domain) Domain
	Login(domain *Domain) (Domain, error)
	// GetAll() ([]Domain, error)
	// GetByID(id int) (Domain, error)
	// Update(domain Domain) (Domain, error)
	// Delete(domain Domain) (Domain, error)
}