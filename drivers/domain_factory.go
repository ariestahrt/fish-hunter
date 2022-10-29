package drivers

import (
	userDomain "fish-hunter/businesses/users"

	userDB "fish-hunter/drivers/mongo/users"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(db *mongo.Database) userDomain.Repository {
	return userDB.NewMongoRepository(db)
}

func NewCronRepository(db *mongo.Database) userDomain.Repository {
	return nil
}