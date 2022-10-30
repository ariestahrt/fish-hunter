package urls

import (
	"context"
	"fish-hunter/businesses/urls"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type urlRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) urls.Repository {
	return &urlRepository{
		collection: db.Collection("urls"),
	}
}

func (u *urlRepository) GetAll() ([]urls.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	
	var urls []Url

	cursor, err := u.collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return ToDomainArray(&urls), err
	}

	if err = cursor.All(ctx, &urls); err != nil {
		return ToDomainArray(&urls), err
	}

	return ToDomainArray(&urls), nil
}

func (u *urlRepository) Save(domain urls.Domain) (urls.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	url := FromDomain(domain)
	url.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	url.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	url.ExecutedStatus = "queued"
	
	// Check for duplicate url
	var urlDomain urls.Domain
	u.collection.FindOne(ctx, map[string]interface{}{
		"url": url.Url,
	}).Decode(&urlDomain)

	if urlDomain.Url != "" {
		return urls.Domain{}, nil
	}
	
	result, err := u.collection.InsertOne(ctx, url)
	if err != nil {
		return urls.Domain{}, err
	}

	url.Id = result.InsertedID.(primitive.ObjectID)
	return url.ToDomain(), nil
}

func (u *urlRepository) GetByID(id string) (urls.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var url Url
	ObjId, _ := primitive.ObjectIDFromHex(id)
	err := u.collection.FindOne(ctx, map[string]interface{}{
		"_id": ObjId,
	}).Decode(&url)

	if err != nil {
		return urls.Domain{}, err
	}

	return url.ToDomain(), nil
}