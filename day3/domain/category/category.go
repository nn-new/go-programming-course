package category

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Category string             `bson:"category" json:"category"`
}
