package catalog

import (
	"context"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCatalog(db *mongo.Database) func(context.Context) ([]Catalog, error) {
	return func(ctx context.Context) ([]Catalog, error) {
		collection := getCatalogCollection(db)

		cur, err := collection.Find(ctx, bson.D{})
		if err != nil {
			log.Error(err)
		}

		results := []Catalog{}
		if err := cur.All(ctx, &results); err != nil {
			log.Error(err)
		}
		cur.Close(ctx)

		return results, err
	}
}
