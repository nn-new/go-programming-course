package category

import (
	"context"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCategories(db *mongo.Database) func(context.Context) ([]Category, error) {
	return func(ctx context.Context) ([]Category, error) {
		collection := getCategoryCollection(db)

		cur, err := collection.Find(ctx, bson.D{})
		if err != nil {
			log.Error(err)
		}

		results := []Category{}
		if err := cur.All(ctx, &results); err != nil {
			log.Error(err)
		}
		cur.Close(ctx)

		return results, err
	}
}
