package users

import (
	"context"
	"errors"
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
	rec.Roles = []string{"guest"}

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
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	
	objId, _ := primitive.ObjectIDFromHex(domain.ID)
	var user User
	
	// Check for duplicate email
	if u.collection.FindOne(ctx, map[string]interface{}{
		"email": domain.Email,
		"_id": map[string]interface{}{
			"$ne": objId,
		},
	}).Decode(&user) != mongo.ErrNoDocuments {
		return users.Domain{}, helpers.ErrDuplicateEmail
	}

	// Check for duplicate username
	if u.collection.FindOne(ctx, map[string]interface{}{
		"username": domain.Username,
		"_id": map[string]interface{}{
			"$ne": objId,
		},
	}).Decode(&user) != mongo.ErrNoDocuments {
		return users.Domain{}, errors.New("username already taken")
	}


	res, err := u.collection.UpdateByID(ctx, objId, map[string]interface{}{
		"$set": map[string]interface{}{
			"username": domain.Username,
			"email": domain.Email,
			"name": domain.Name,
			"phone": domain.Phone,
			"university": domain.University,
			"position": domain.Position,
			"proposal": domain.Proposal,
			"updated_at": time.Now(),
		},
	})

	if res.ModifiedCount == 0 {
		return users.Domain{}, errors.New("failed to update profile")
	}

	if err != nil {
		return users.Domain{}, err
	}

	u.collection.FindOne(ctx, map[string]interface{}{
		"_id": objId,
	}).Decode(&user)
	
	return user.ToDomain(), nil
}

func (u *userRepository) UpdatePassword(domain *users.Domain) (users.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(domain.ID)
	var user User

	// Validate old password
	u.collection.FindOne(ctx, map[string]interface{}{
		"_id": objId,
	}).Decode(&user)
	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(domain.Password)); err != nil {
		return users.Domain{}, errors.New("old password is incorrect")
	}

	// Update password
	password, _ := bcrypt.GenerateFromPassword([]byte(domain.NewPassword), bcrypt.MinCost)

	res, err := u.collection.UpdateByID(ctx, objId, map[string]interface{}{
		"$set": map[string]interface{}{
			"password": string(password),
			"updated_at": time.Now(),
		},
	})

	if res.ModifiedCount == 0 {
		return users.Domain{}, errors.New("failed to update password")
	}

	if err != nil {
		return users.Domain{}, err
	}

	u.collection.FindOne(ctx, map[string]interface{}{
		"_id": objId,
	}).Decode(&user)

	return user.ToDomain(), nil
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

func (u *userRepository) GetAllUsers() ([]users.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var users []User

	cursor, err := u.collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return ToDomainArray(&users), err
	}

	if err = cursor.All(ctx, &users); err != nil {
		return ToDomainArray(&users), err
	}

	return ToDomainArray(&users), nil
}

func (u *userRepository) GetByID(id string) (users.Domain, error) {
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

func (u *userRepository) Update(domain *users.Domain) (users.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(domain.ID)
	var user User

	// Check for duplicate email
	if u.collection.FindOne(ctx, map[string]interface{}{
		"email": domain.Email,
		"_id": map[string]interface{}{
			"$ne": objId,
		},
	}).Decode(&user) != mongo.ErrNoDocuments {
		return users.Domain{}, helpers.ErrDuplicateEmail
	}

	// Check for duplicate username
	if u.collection.FindOne(ctx, map[string]interface{}{
		"username": domain.Username,
		"_id": map[string]interface{}{
			"$ne": objId,
		},
	}).Decode(&user) != mongo.ErrNoDocuments {
		return users.Domain{}, errors.New("username already taken")
	}

	// Get Old Data
	u.collection.FindOne(ctx, map[string]interface{}{
		"_id": objId,
	}).Decode(&user)

	// check if password is empty
	if domain.Password == "" {
		domain.Password = user.Password
	} else {
		// new password
		password, _ := bcrypt.GenerateFromPassword([]byte(domain.Password), bcrypt.MinCost)
		domain.Password = string(password)
	}

	res, err := u.collection.UpdateByID(ctx, objId, map[string]interface{}{
		"$set": map[string]interface{}{
			"username": domain.Username,
			"email": domain.Email,
			"password": domain.Password,
			"is_active": domain.IsActive,
			"name": domain.Name,
			"phone": domain.Phone,
			"university": domain.University,
			"position": domain.Position,
			"proposal": domain.Proposal,
			"roles": domain.Roles,
			"updated_at": time.Now(),
		},
	})

	if res.ModifiedCount == 0 {
		return users.Domain{}, errors.New("failed to update profile")
	}

	if err != nil {
		return users.Domain{}, err
	}

	u.collection.FindOne(ctx, map[string]interface{}{
		"_id": objId,
	}).Decode(&user)
	
	return user.ToDomain(), nil
}