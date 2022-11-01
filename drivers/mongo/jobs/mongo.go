package jobs

import (
	"context"
	"fish-hunter/businesses/jobs"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type jobRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) jobs.Repository {
	return &jobRepository{
		collection: db.Collection("jobs"),
	}
}

func (u *jobRepository) GetAll() ([]jobs.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var jobs []Job

    cursor, err := u.collection.Aggregate(ctx, bson.A{
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "urls"},
					{Key: "localField", Value: "ref_url"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "urls"},
				},
			},
		},
		bson.D{
			{Key: "$set",
				Value: bson.D{
					{Key: "url",
						Value: bson.D{
							{Key: "$arrayElemAt",
								Value: bson.A{
									"$urls.url",
									0,
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{Key: "$project",
				Value: bson.D{
					{Key: "urls", Value: 0},
					{Key: "ref_url", Value: 0},
				},
			},
		},
	})

	if err != nil {
		return ToDomainArray(&jobs), err
	}

	if err = cursor.All(ctx, &jobs); err != nil {
		return ToDomainArray(&jobs), err
	}

	return ToDomainArray(&jobs), nil
}

func (u *jobRepository) GetByID(id string) (jobs.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var job Job
	ObjId,_ := primitive.ObjectIDFromHex(id)

	err := u.collection.FindOne(ctx, map[string]interface{}{
		"_id": ObjId,
	}).Decode(&job)
	if err != nil {
		return jobs.Domain{}, err
	}

	return job.ToDomain(), nil
}

// Count Total
func (u *jobRepository) CountTotal() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	count, err := u.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Get Total Between Dates
func (u *jobRepository) GetTotalBetweenDates(startDate, endDate time.Time) (int64, error) {
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