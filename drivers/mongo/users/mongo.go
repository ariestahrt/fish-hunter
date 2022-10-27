package users

import (
	"context"
	"fish-hunter/businesses/users"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) users.Repository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

func (u *userRepository) Register(domain *users.Domain) users.Domain {
	// check for duplicate email
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var user User

	err := u.collection.FindOne(ctx, map[string]interface{}{
		"email": domain.Email,
	}).Decode(&user)

	if err == nil {
		return users.Domain{}
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(domain.Password), bcrypt.MinCost)
	rec := FromDomain(domain)
	rec.Password = string(password)
	rec.CreatedAt = time.Now()
	rec.UpdatedAt = time.Now()

	defer cancel()

	u.collection.InsertOne(ctx, rec)

	return rec.ToDomain()
}

func (u *userRepository) Login(domain *users.Domain) (users.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var user User

	err := u.collection.FindOne(ctx, map[string]interface{}{
		"email": domain.Email,
	}).Decode(&user)
	
	if err != nil {
		return users.Domain{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(domain.Password))
	if err != nil {
		return users.Domain{}, err
	}

	return *domain, nil
}