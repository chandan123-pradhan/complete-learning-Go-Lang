package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Id      primitive.ObjectID `json:"_id"`
	Movie   string             `json:"movie"`
	Watched bool               `json:"watched"` //it mean comma and , not allowed
}
