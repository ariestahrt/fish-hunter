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
					{Key: "created_at", Value: 0},
					{Key: "updated_at", Value: 0},
					{Key: "deleted_at", Value: 0},
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

	// Update status and screenshot path
	_, err := u.collection.UpdateOne(ctx, bson.M{"_id": domain.Id}, bson.M{"$set": bson.M{"status": domain.Status, "screenshot_path": domain.ScreenshotPath}})

	if err != nil {
		return datasets.Domain{}, err
	}

	// Get updated data
	var dataset Dataset
	u.collection.FindOne(ctx, bson.M{"_id": domain.Id}).Decode(&dataset)

	return dataset.ToDomain(), nil
}

func (u *datasetRepository) TopBrands() (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cursor, _ := u.collection.Aggregate(ctx, bson.A{
		bson.M{"$match": bson.M{"$or": []bson.M{{"status": "ok"}, {"status": "new"}}}},
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

// Count Total Valid
func (u *datasetRepository) CountTotalValid() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	count, _ := u.collection.CountDocuments(ctx, bson.M{"status": "ok"})

	return count, nil
}

// Count Total All
func (u *datasetRepository) CountTotal() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	count, _ := u.collection.CountDocuments(ctx, bson.M{})

	return count, nil
}

// Get Total Between Dates
func (u *datasetRepository) GetTotalBetweenDates(startDate time.Time, endDate time.Time) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	count, _ := u.collection.CountDocuments(ctx, bson.M{"created_at": bson.M{"$gte": startDate, "$lte": endDate}})

	return count, nil
}