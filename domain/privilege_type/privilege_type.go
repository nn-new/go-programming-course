package privilegetype 

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PrivilegeType struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type string             `bson:"type" json:"type"`
}
