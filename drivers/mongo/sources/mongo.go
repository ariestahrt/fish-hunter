package sources

import (
	"context"
	"fish-hunter/businesses/sources"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type sourceRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) sources.Repository {
	return &sourceRepository{
		collection: db.Collection("sources"),
	}
}

func (u *sourceRepository) GetByName(name string) (sources.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	
	var source Source

	if err := u.collection.FindOne(ctx, map[string]interface{}{
		"name": name,
	}).Decode(&source); err != nil {
		return source.ToDomain(), err
	}

	return source.ToDomain(), nil
}