package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type UniversityInfo struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	University_id int                `json:"university_id"`
	Domain        string             `json:"domain"`
	Web_page      string             `json:"web_page"`
}
