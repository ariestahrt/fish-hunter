package helpers

import (
	"fish-hunter/businesses/sources"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSourceInformation(source string) sources.Domain {
	switch source {
	case "openphish":
		objId,_ := primitive.ObjectIDFromHex("6324ab1347bb2e854efbe6d2")
		return sources.Domain{
			Name: "OpenPhish",
			Url:  "https://openphish.com/",
			Id:   objId,
		}
	default:
		return sources.Domain{}
	}
}