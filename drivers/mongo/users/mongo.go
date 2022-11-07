package users

import (
	"context"
	"fish-hunter/businesses/users"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) users.Repository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

// Create
func (u *userRepository) Create(domain *users.Domain) (users.Domain, error) {
	ctx, cancell := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancell()

	// Insert to mongo
	res, err := u.collection.InsertOne(ctx, FromDomain(domain))
	if err != nil {
		return users.Domain{}, err
	}
	
	// Get inserted data
	var user User
	err = u.collection.FindOne(ctx, map[string]interface{}{
		"_id": res.InsertedID,
	}).Decode(&user)

	return user.ToDomain(), err
}

// Get By Email
func (u *userRepository) GetByEmail(email string) (users.Domain, error) {
	ctx, cancell := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancell()

	var user User

	// Get user by email and not deleted
	err := u.collection.FindOne(ctx, map[string]interface{}{
		"email": email,
		"deleted_at": map[string]interface{}{
			"$eq": primitive.NilObjectID.Timestamp(),
		},
	}).Decode(&user)

	return user.ToDomain(), err
}

// Get By Username
func (u *userRepository) GetByUsername(username string) (users.Domain, error) {
	ctx, cancell := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancell()

	var user User

	// Get user by username and not deleted
	err := u.collection.FindOne(ctx, map[string]interface{}{
		"username": username,
		"deleted_at": map[string]interface{}{
			"$eq": primitive.NilObjectID.Timestamp(),
		},
	}).Decode(&user)

	return user.ToDomain(), err
}

// Update
func (u *userRepository) Update(old *users.Domain, new *users.Domain) (users.Domain, error) {
	ctx, cancell := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancell()

	// Update to mongo
	_, err := u.collection.UpdateOne(ctx, map[string]interface{}{
		"_id": old.Id,
	}, map[string]interface{}{
		"$set": FromDomain(new),
	})
	if err != nil {
		return users.Domain{}, err
	}

	// Get updated data
	var user User
	err = u.collection.FindOne(ctx, map[string]interface{}{
		"_id": old.Id,
	}).Decode(&user)

	return user.ToDomain(), err
}

// Delete
func (u *userRepository) Delete(id primitive.ObjectID) (users.Domain, error) {
	ctx, cancell := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancell()

	// Delete user
	_, err := u.collection.UpdateOne(ctx, map[string]interface{}{
		"_id": id,
	}, map[string]interface{}{
		"$set": map[string]interface{}{
			"deleted_at": primitive.NewDateTimeFromTime(time.Now()),
		},
	})
	if err != nil {
		return users.Domain{}, err
	}

	// Get deleted data
	var user User
	err = u.collection.FindOne(ctx, map[string]interface{}{
		"_id": id,
	}).Decode(&user)

	return user.ToDomain(), err
}

// GetByID
func (u *userRepository) GetByID(id primitive.ObjectID) (users.Domain, error) {
	ctx, cancell := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancell()

	var user User

	// Get user by id
	err := u.collection.FindOne(ctx, map[string]interface{}{
		"_id": id,
	}).Decode(&user)

	return user.ToDomain(), err
}

// Get All
func (u *userRepository) GetAll() ([]users.Domain, error) {
	ctx, cancell := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancell()

	var users []User

	// Get all users
	cursor, err := u.collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return ToDomainArray(&users), err
	}

	// Decode cursor to users
	err = cursor.All(ctx, &users)
	if err != nil {
		return ToDomainArray(&users), err
	}

	return ToDomainArray(&users), err
}