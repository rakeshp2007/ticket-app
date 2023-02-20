package controllers

import (
	"context"
	"fmt"
	"net/http"
	database "ticket-app/internal/app/utils/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Contact struct {
	Id       string `bson:"id" json:"id" validate:"required"`
	Name     string `bson:"name" json:"name" validate:"required"`
	LastName string `bson:"lastName" json:"lastName" validate:"required"`
	Email    string `bson:"email" json:"email" validate:"required"`
}
type Assignee struct {
	Id       string `bson:"id" json:"id" validate:"required"`
	Name     string `bson:"name" json:"name" validate:"required"`
	LastName string `bson:"lastName" json:"lastName" validate:"required"`
	Email    string `bson:"email" json:"email" validate:"required"`
}

func EditTicket(c *gin.Context) {

	validateStat, validateMessage, docId, _ := ValidateId(c)
	if !validateStat {
		c.JSON(http.StatusBadRequest, gin.H{"eror": map[string]interface{}{"Contact": validateMessage}})
		return
	}

	validateStat, TicketCreateValidate := ValidateTicket(c)
	if !validateStat {
		return
	}
	validateStat, validateMessage, ContactDetails := ValidateContact(fmt.Sprint(TicketCreateValidate.Contact))
	if !validateStat {
		c.JSON(http.StatusBadRequest, gin.H{"eror": map[string]interface{}{"Contact": validateMessage}})
		return
	}

	validateStat, validateMessage, UserDetails := validateUser(fmt.Sprint(TicketCreateValidate.Assignee))
	if !validateStat {
		c.JSON(http.StatusBadRequest, gin.H{"eror": map[string]interface{}{"Contact": validateMessage}})
		return
	}

	ticketSaveData := GetTicketSaveData(true, TicketCreateValidate, ContactDetails, UserDetails)
	bsonSaveData, err := ToDoc(ticketSaveData)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//filter := bson.D{{"species", "Ledebouria socialis"}, {"plant_id", 3}}
	ticketCollections := database.GetCollection(database.MongoDB, "tickets")
	filter := bson.D{{Key: "_id", Value: docId}}
	update := bson.D{{Key: "$set", Value: bsonSaveData}}
	opts := options.Update().SetUpsert(true)
	//fmt.Println(bsonSaveData)
	_, err = ticketCollections.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ticket updated successfully"})
	return

}
