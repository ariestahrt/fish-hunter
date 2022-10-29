package drivers

import (
	sourceDomain "fish-hunter/businesses/sources"
	urlDomain "fish-hunter/businesses/urls"
	userDomain "fish-hunter/businesses/users"

	sourceDB "fish-hunter/drivers/mongo/sources"
	urlDB "fish-hunter/drivers/mongo/urls"
	userDB "fish-hunter/drivers/mongo/users"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(db *mongo.Database) userDomain.Repository {
	return userDB.NewMongoRepository(db)
}

func NewCronRepository(db *mongo.Database) userDomain.Repository {
	return nil
}

func NewUrlRepository(db *mongo.Database) urlDomain.Repository {
	return urlDB.NewMongoRepository(db)
}

func NewSourceRepository(db *mongo.Database) sourceDomain.Repository {
	return sourceDB.NewMongoRepository(db)
}