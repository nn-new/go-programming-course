package privilege

import (
	"context"
	"privilege/domain/pagination"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPrivilege(db *mongo.Database) func(context.Context, pagination.Pagination) ([]Privilege, error) {
	return func(ctx context.Context, pag pagination.Pagination) ([]Privilege, error) {
		collection := getPrivilegeCollection(db)

		var filterTitle, filterPrivilegeId, filterDiscountText, filterSearch primitive.E
		if pag.Search != "" {
			filterTitle = buildFilterSearch("title", pag.Search)
			filterPrivilegeId = buildFilterSearch("privilege_id", pag.Search)
			filterDiscountText = buildFilterSearch("discount_text", pag.Search)
			filterSearch = bson.E{
				Key: "$or", Value: []bson.D{
					{filterTitle},
					{filterPrivilegeId},
					{filterDiscountText},
				},
			}
		}

		matchStage := bson.D{{
			Key: "$match", Value: bson.D{
				filterSearch,
				{Key: "is_deleted", Value: nil},
			},
		}}

		limitStage := bson.D{{
			Key: "$limit", Value: pag.GetPageSize(),
		}}
		skipStage := bson.D{{
			Key: "$skip", Value: (pag.GetPage() - 1) * pag.GetPageSize(),
		}}

		var sortStage primitive.D
		if pag.Sort != "" {
			sortStage = bson.D{{Key: "$sort", Value: bson.D{{Key: pag.Sort, Value: pag.GetDirection()}}}}
		} else {
			sortStage = bson.D{{Key: "$sort", Value: bson.D{{Key: "_id", Value: pag.GetDirection()}}}}
		}

		pipeline := mongo.Pipeline{matchStage, sortStage, skipStage, limitStage}

		cur, err := collection.Aggregate(ctx, pipeline)

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

func buildFilterSearch(key, searchKeyword string) bson.E {
	return bson.E{
		Key: key, Value: primitive.Regex{
			Pattern: searchKeyword,
			Options: "im",
		},
	}
}
