package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRetrieve struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id" validate:"required"`
	UserName    string             `bson:"userName" json:"userName,omitempty" validate:"required"`
	Name        string             `bson:"name" json:"name,omitempty" validate:"required"`
	LastName    string             `bson:"lastName" json:"lastName,omitempty" validate:"required"`
	Password    string             `bson:"password" json:"password,omitempty" validate:"required"`
	CreatedDate primitive.DateTime `bson:"createdDate" json:"createdDate,omitempty"`
	Status      string             `bson:"status" json:"status,omitempty" validate:"required"`
	Roles       []string           `bson:"roles" json:"roles,omitempty"`
	TimeZone    string             `bson:"timezone" json:"timezone,omitempty"`
}
