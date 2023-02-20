package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactRetrieve struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"_id" validate:"required"`
	Name           string             `json:"name,omitempty" validate:"required"`
	LastName       string             `json:"lastName,omitempty" validate:"required"`
	Email          string             `json:"email,omitempty" validate:"required"`
	CreatedDate    primitive.DateTime `json:"createdDate,omitempty"`
	CompanyDetails struct {
		Name     string `json:"name,omitempty"`
		Location string `json:"location,omitempty"`
		Pin      int32  `json:"pin,omitempty"`
	} `json:"companyDetails,omitempty"`
	Phone []string `json:"phone,omitempty"`
}
