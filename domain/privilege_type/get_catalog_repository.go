package privilegetype

import (
	"context"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPrivilegeTypes(db *mongo.Database) func(context.Context) ([]PrivilegeType, error) {
	return func(ctx context.Context) ([]PrivilegeType, error) {
		collection := getPrivilegeTypeCollection(db)

		cur, err := collection.Find(ctx, bson.D{})
		if err != nil {
			log.Error(err)
		}

		results := []PrivilegeType{}
		if err := cur.All(ctx, &results); err != nil {
			log.Error(err)
		}
		cur.Close(ctx)

		return results, err
	}
}
