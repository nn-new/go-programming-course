package catalog

import "go.mongodb.org/mongo-driver/mongo"

func getCatalogCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("catalogs")
}
