package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Domain struct {
	Id        primitive.ObjectID    `bson:"_id,omitempty" json:"_id"`
	Username  string				`bson:"username" json:"username"`
	Email     string				`bson:"email" json:"email"`
	Password  string				`bson:"password" json:"password"`
	IsActive  bool					`bson:"is_active" json:"is_active"`
	Name	  string				`bson:"name" json:"name"`
	Phone	  string 			  	`bson:"phone" json:"phone"`
	University string 			  	`bson:"university" json:"university"`
	Position   string 				`bson:"position" json:"position"` // Student, Lecturer, Staff 
	Proposal   string				`bson:"proposal" json:"proposal"`
	Roles     []string		  		`bson:"roles" json:"roles"`
	CreatedAt primitive.DateTime	`bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime	`bson:"updated_at" json:"updated_at"`
	DeletedAt primitive.DateTime	`bson:"deleted_at" json:"deleted_at"`
}

type UseCase interface {
	Register(domain *Domain) (Domain, error)
	Login(domain *Domain) (string, error)
	UpdateProfile(old *Domain, new *Domain) (Domain, error)
	UpdatePassword(old *Domain, new *Domain) (Domain, error)
	GetProfile(id string) (Domain, error)
	GetAllUsers() ([]Domain, error)
	GetByID(id string) (Domain, error)
	UpdateByAdmin(new *Domain) (Domain, error)
	Delete(id string) (Domain, error)
}

type Repository interface {
	// New
	Create(domain *Domain) (Domain, error)
	Update(old *Domain, new *Domain) (Domain, error)
	Delete(id primitive.ObjectID) (Domain, error)
	GetByID(id primitive.ObjectID) (Domain, error)
	GetByEmail(email string) (Domain, error)
	GetByUsername(username string) (Domain, error)
	GetAll() ([]Domain, error)
}