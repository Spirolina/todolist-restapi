package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Title    string             `bson:"title" json:"title"`
	Complete bool               `bson:"complete" json:"complete"`
}
