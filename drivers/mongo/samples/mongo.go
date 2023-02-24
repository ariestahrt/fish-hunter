package samples

import (
	"context"
	"fish-hunter/businesses/samples"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type sampleRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) samples.Repository {
	return &sampleRepository{
		collection: db.Collection("samples"),
	}
}

func (u *sampleRepository) GetAll() ([]samples.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var samples []Sample

    cursor, err := u.collection.Find(ctx, bson.M{})

	if err != nil {
		return ToDomainArray(&samples), err
	}

	if err = cursor.All(ctx, &samples); err != nil {
		return ToDomainArray(&samples), err
	}

	return ToDomainArray(&samples), nil
}

func (u *sampleRepository) GetByID(id string) (samples.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var sample Sample
	ObjId,_ := primitive.ObjectIDFromHex(id)

	err := u.collection.FindOne(ctx, map[string]interface{}{
		"_id": ObjId,
	}).Decode(&sample)
	if err != nil {
		return samples.Domain{}, err
	}

	return sample.ToDomain(), nil
}

// Count Total
func (u *sampleRepository) CountTotal() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	count, err := u.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Get Total Between Dates
func (u *sampleRepository) GetTotalBetweenDates(startDate, endDate time.Time) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	count, err := u.collection.CountDocuments(ctx, bson.M{
		"created_at": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	})
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Update
func (u *sampleRepository) Update(id string, sample *samples.Domain) (samples.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	ObjId,_ := primitive.ObjectIDFromHex(id)

	_, err := u.collection.UpdateOne(ctx, bson.M{
		"_id": ObjId,
	}, bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
			"brands": sample.Brands,
			"language": sample.Language,
			"details": sample.Details,
			"type": sample.Type,
		},
	})

	if err != nil {
		return samples.Domain{}, err
	}

	return *sample, nil
}