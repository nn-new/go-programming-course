package privilege

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdatePrivilege(db *mongo.Database) func(context.Context, Privilege) error {
	return func(ctx context.Context, privilege Privilege) error {
		collection := getPrivilegeCollection(db)

		filter := bson.M{"$and": []bson.M{
			{"_id": bson.M{"$eq": privilege.ID}},
			{"is_deleted": bson.M{"$ne": true}},
		}}

		ur, err := collection.ReplaceOne(context.Background(), filter, privilege)
		if ur.ModifiedCount == 0 {
			return errors.New("privilege can not update")
		}

		return err
	}
}
