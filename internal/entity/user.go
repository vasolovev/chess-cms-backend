package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name        string               `json:"name" bson:"name,omitempty"`
	Surname     string               `json:"surname" bson:"surname,omitempty"`
	Patronymic  string               `json:"patronymic" bson:"patronymic,omitempty"`
	Email       string               `json:"email" bson:"email,omitempty"`
	TelegramID  string               `json:"telegramID" bson:"telegramID,omitempty"`
	GroupNumber string               `json:"groupNumber" bson:"groupNumber,omitempty"`
	Lichess     []primitive.ObjectID `json:"lichess" bson:"lichess"`
}
