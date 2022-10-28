package users

import (
	"context"
	"fish-hunter/businesses/users"
	"fish-hunter/helpers"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (u *userRepository) Register(domain *users.Domain) (users.Domain, error) {
	// check for duplicate email
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var user User

	if u.collection.FindOne(ctx, map[string]interface{}{
		"email": domain.Email,
	}).Decode(&user) != mongo.ErrNoDocuments {
		return users.Domain{}, helpers.ErrDuplicateEmail
	}
	
	password, _ := bcrypt.GenerateFromPassword([]byte(domain.Password), bcrypt.MinCost)
	rec := FromDomain(domain)
	rec.Password = string(password)
	rec.IsActive = false
	rec.CreatedAt = time.Now()
	rec.UpdatedAt = time.Now()

	// Force change role to user
	rec.Roles = []string{"user"}

	defer cancel()

	u.collection.InsertOne(ctx, rec)

	return rec.ToDomain(), nil
}

func (u *userRepository) Login(domain *users.Domain) (users.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var user User

	err := u.collection.FindOne(ctx, map[string]interface{}{
		"email": domain.Email,
	}).Decode(&user)

	fmt.Println(user)
	
	if err != nil {
		return users.Domain{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(domain.Password))
	if err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}

func (u *userRepository) UpdateProfile(domain *users.Domain) (users.Domain, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

	// var user User


	panic("implement me")
}

func (u *userRepository) UpdatePassword(domain *users.Domain) (users.Domain, error) {
	panic("implement me")
}

func (u *userRepository) GetProfile(id string) (users.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var user User

	objId, _ := primitive.ObjectIDFromHex(id)

	err := u.collection.FindOne(ctx, map[string]interface{}{
		"_id": objId,
	}).Decode(&user)

	if err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}