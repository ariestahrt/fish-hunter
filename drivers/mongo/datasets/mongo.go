package datasets

import (
	"context"
	"fish-hunter/businesses/datasets"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type datasetRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) datasets.Repository {
	return &datasetRepository{
		collection: db.Collection("datasets"),
	}
}

func (u *datasetRepository) Status(status string) ([]datasets.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var datasets []Dataset
	
	cursor, err := u.collection.Aggregate(ctx, bson.A{
        bson.D{{Key: "$match", Value: bson.D{{Key: "status", Value: status}}}},
        bson.D{
			{Key: "$project",
				Value: bson.D{
					{Key: "ref_url", Value: 0},
					{Key: "ref_job", Value: 0},
					{Key: "domain", Value: 0},
					{Key: "dataset_path", Value: 0},
					{Key: "created", Value: 0},
					{Key: "updated", Value: 0},
				},
			},
		},
	})

	if err != nil {
		return ToDomainArray(&datasets), err
	}

	if err = cursor.All(ctx, &datasets); err != nil {
		return ToDomainArray(&datasets), err
	}

	return ToDomainArray(&datasets), nil
}

func (u *datasetRepository) GetByID(id string) (datasets.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var dataset Dataset

	ObjId,_ := primitive.ObjectIDFromHex(id)
	err := u.collection.FindOne(ctx, bson.M{"_id": ObjId}).Decode(&dataset)
	if err != nil {
		return dataset.ToDomain(), err
	}

	return dataset.ToDomain(), nil
}

func (u *datasetRepository) Validate(domain datasets.Domain) (datasets.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := u.collection.UpdateOne(ctx, bson.M{"_id": domain.Id}, bson.M{"$set": bson.M{"status": domain.Status}})

	if err != nil {
		return datasets.Domain{}, err
	}

	return domain, nil
}

func (u *datasetRepository) TopBrands() (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cursor, _ := u.collection.Aggregate(ctx, bson.A{
		bson.D{{Key: "$project", Value: bson.D{{Key: "brands", Value: 1}}}},
	})

	brands := make(map[string]interface{})

	for cursor.Next(ctx) {
		var singleDataset Dataset
		cursor.Decode(&singleDataset)

		for _, brand := range singleDataset.Brands {
			if _, ok := brands[brand]; ok {
				brands[brand] = brands[brand].(int) + 1
			} else {
				brands[brand] = 1
			}
		}
	}

	return brands, nil
}