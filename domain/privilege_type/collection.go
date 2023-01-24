package privilegetype

import "go.mongodb.org/mongo-driver/mongo"

func getPrivilegeTypeCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("privilege_type")
}
