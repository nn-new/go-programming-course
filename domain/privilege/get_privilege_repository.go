package privilege

import (
	"context"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPrivilege(db *mongo.Database) func(context.Context) ([]Privilege, error) {
	return func(ctx context.Context) ([]Privilege, error) {
		collection := getPrivilegeCollection(db)

		filter := bson.M{"is_deleted": bson.M{"$ne": true}}

		cur, err := collection.Find(ctx, filter)
		if err != nil {
			log.Error(err)
		}

		// limitStage := bson.D{{
		// 	Key: "$limit", Value: pag.GetPageSize(),
		// }}
		// skipStage := bson.D{{
		// 	Key: "$skip", Value: (pag.GetPage() - 1) * pag.GetPageSize(),
		// }}

		// var sortStage primitive.D
		// if pag.Sort != "" {
		// 	sortStage = bson.D{{Key: "$sort", Value: bson.D{
		// 		{Key: pag.Sort, Value: pag.GetDirection()},
		// 	},
		// 	}}
		// } else {
		// 	sortStage = bson.D{{Key: "$sort", Value: bson.D{{Key: "_id", Value: pag.GetDirection()}}}}
		// }

		// pipeline := mongo.Pipeline{filter, sortStage, skipStage, limitStage}

		results := []Privilege{}
		if err := cur.All(ctx, &results); err != nil {
			log.Error(err)
		}
		cur.Close(ctx)

		return results, err
	}
}

func GetPrivilegeById(db *mongo.Database) func(context.Context, primitive.ObjectID) (Privilege, error) {
	return func(ctx context.Context, id primitive.ObjectID) (Privilege, error) {
		collection := getPrivilegeCollection(db)

		filter := bson.M{"$and": []bson.M{
			{"_id": bson.M{"$eq": id}},
			{"is_deleted": bson.M{"$ne": true}},
		}}


		privilege := Privilege{}
		err := collection.FindOne(ctx, filter).Decode(&privilege)
		return privilege, err
	}
}
