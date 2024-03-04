package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitemtpy" bson: "_id,omitemtpy"`
	Movie   string             `json:"movie,omitemtpy"`
	Watched bool               `json:"watched,omitemtpy"` //it mean comma and , not allowed
}
