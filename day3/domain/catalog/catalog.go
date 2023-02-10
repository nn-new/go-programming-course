package catalog

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Catalog struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Catalog string             `bson:"catalog" json:"catalog"`
}
