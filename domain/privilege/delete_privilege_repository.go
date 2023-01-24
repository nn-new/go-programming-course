package privilege

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeletePrivilege(db *mongo.Database) func(context.Context, primitive.ObjectID) error {
	return func(ctx context.Context, id primitive.ObjectID) error {
		collection := getPrivilegeCollection(db)

		filter := bson.M{"_id": id}

		ur, err := collection.UpdateOne(context.Background(), filter, bson.M{
			"$set": bson.M{"is_deleted": true},
		})

		if ur.ModifiedCount == 0 {
			return errors.New("privilege can not delete")
		}

		return err
	}
}
