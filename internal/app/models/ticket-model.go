package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateTicketValidate struct {
	Subject     string `json:"subject" validate:"required,max=1000"`
	Description string `json:"description" validate:"required,max=2000"`
	Status      string `json:"status" validate:"required,oneof=open pending resolved closed"`
	Priority    string `json:"priority" validate:"required,oneof=low medium high"`
	Contact     string `json:"contact" validate:"required"`
	Assignee    string `json:"assignee" validate:"required"`
}

type (
	TicketJson struct {
		ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		Subject     string             `bson:"subject" json:"subject" validate:"required,max=1000"`
		Description string             `bson:"description" json:"description" validate:"required,max=2000"`
		Status      string             `bson:"status" json:"status" validate:"required,oneof=open pending resolved closed"`
		Priority    string             `bson:"priority" json:"priority" validate:"required,oneof=low medium high"`
		CreatedDate primitive.DateTime `bson:"createdDate,omitempty" json:"-"`
		DateCreated string             `bson:"-,omitempty" json:"createdDate,omitempty"`
		Contact     struct {
			Id       string `bson:"id" json:"id" validate:"required"`
			Name     string `bson:"name" json:"name" validate:"required"`
			LastName string `bson:"lastName" json:"lastName" validate:"required"`
			Email    string `bson:"email" json:"email" validate:"required"`
		} `bson:"contact" json:"contact" validate:"required"`
		Assignee struct {
			Id       string `bson:"id" json:"id" validate:"required"`
			Name     string `bson:"name" json:"name" validate:"required"`
			LastName string `bson:"lastName" json:"lastName" validate:"required"`
			Email    string `bson:"email" json:"email" validate:"required"`
		} `bson:"assignee" json:"assignee" validate:"required"`
	}
)

type TicketListRequest struct {
	Page      int    `json:"page" validate:"required,numeric,min=1"`
	Limit     int    `json:"limit" validate:"required,numeric,min=1,max=10000"`
	DateFrom  string `json:"date_from"  validate:"omitempty,datetime=2006-01-02" `
	DateTo    string `json:"date_to"  validate:"omitempty,datetime=2006-01-02" `
	Status    string `json:"status" validate:"omitempty,oneof=open pending resolved closed"`
	Priority  string `json:"priority" validate:"omitempty,oneof=low medium high"`
	Search    string `json:"search" validate:"omitempty"`
	SortField string `json:"sort_field" validate:"omitempty,oneof=createdDate subject"`
	SortOrder string `json:"sort_order" validate:"omitempty,oneof=asc desc"`
}

type (
	TicketChange struct {
		Status   string `bson:"status,omitempty" json:"status,omitempty" validate:"omitempty,oneof=open pending resolved closed"`
		Priority string `bson:"priority,omitempty" json:"priority,omitempty" validate:"omitempty,oneof=low medium high"`
	}
)
