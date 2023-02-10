package privilege

import "go.mongodb.org/mongo-driver/mongo"

func getPrivilegeCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("privileges")
}
