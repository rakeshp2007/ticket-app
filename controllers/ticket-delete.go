package controllers

import (
	"context"
	"net/http"
	sr "ticket-app/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteTicket(c *gin.Context) {

	validateStat, validateMessage, docId, _ := ValidateId(c)
	if !validateStat {
		c.JSON(http.StatusBadRequest, gin.H{"eror": map[string]interface{}{"Contact": validateMessage}})
		return
	}
	ticketCollections := sr.GetCollection(sr.MongoDB, "tickets")

	result, err := ticketCollections.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: docId}})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result.DeletedCount > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
		return
	}
}
