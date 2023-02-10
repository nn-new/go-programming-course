package privilege

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePrivilege(db *mongo.Database) func(context.Context, Privilege) error {
	return func(ctx context.Context, privilege Privilege) error {
		collection := getPrivilegeCollection(db)
		_, err := collection.InsertOne(ctx, privilege)
		return err
	}
}
