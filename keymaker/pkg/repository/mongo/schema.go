package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Key ...
type Key struct {
	ID   primitive.ObjectID `bson:"_id"`
	Key  string             `bson:"key"`
	Used bool               `bson:"used"`
}
