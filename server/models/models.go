package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDo struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Tasks  string             `json:"tasks,omitempty"`
	Status bool               `json:"status,omitempty"`
}
