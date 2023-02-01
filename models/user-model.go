package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRetrieve struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id" validate:"required"`
	UserName    string             `bson:"userName" json:"userName" validate:"required"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	LastName    string             `bson:"lastName" json:"lastName" validate:"required"`
	Password    string             `bson:"password" json:"password" validate:"required"`
	CreatedDate primitive.DateTime `bson:"createdDate" json:"createdDate"`
	Status      string             `bson:"status" json:"status" validate:"required"`
	Roles       []string           `bson:"roles" json:"roles"`
	TimeZone    string             `bson:"timezone" json:"timezone"`
}
