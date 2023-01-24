package category

import "go.mongodb.org/mongo-driver/mongo"

func getCategoryCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("categories")
}
