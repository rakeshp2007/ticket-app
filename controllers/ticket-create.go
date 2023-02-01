package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	sr "ticket-app/services"

	"github.com/gin-gonic/gin"
)

func CreateTicket(c *gin.Context) {
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
	ticketSaveData := GetTicketSaveData(false, TicketCreateValidate, ContactDetails, UserDetails)
	ticketCollections := sr.GetCollection(sr.MongoDB, "tickets")
	result, err := ticketCollections.InsertOne(context.TODO(), ticketSaveData)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
	return
}
