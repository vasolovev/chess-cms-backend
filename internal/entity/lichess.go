package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lichess struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Nickname string             `json:"nickname" bson:"nickname"`
	Ban      bool               `json:"ban" bson:"ban"`
}
