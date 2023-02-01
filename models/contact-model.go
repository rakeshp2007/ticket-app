package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactRetrieve struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"_id" validate:"required"`
	Name           string             `json:"name" validate:"required"`
	LastName       string             `json:"lastName" validate:"required"`
	Email          string             `json:"email" validate:"required"`
	CreatedDate    primitive.DateTime `json:"createdDate"`
	CompanyDetails struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Pin      int32  `json:"pin"`
	} `json:"companyDetails"`
	Phone []string `json:"phone"`
}
