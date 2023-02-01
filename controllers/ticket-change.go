package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	sr "ticket-app/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ChangeTicket(c *gin.Context) {
	url := strings.Split(fmt.Sprint(c.Request.URL.Path), "/")
	statusType := (url[len(url)-2])
	validateStat, validateMessage, docId, _ := ValidateId(c)
	if !validateStat {
		c.JSON(http.StatusBadRequest, gin.H{"eror": map[string]interface{}{"Contact": validateMessage}})
		return
	}
	validateStat, TicketChangeValidate := ValidateTicketChange(c)
	if !validateStat {
		return
	}
	ticketSaveData := GetTicketChangeData(statusType, TicketChangeValidate)
	bsonSaveData, err := ToDoc(ticketSaveData)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//filter := bson.D{{"species", "Ledebouria socialis"}, {"plant_id", 3}}
	fmt.Println(bsonSaveData)
	ticketCollections := sr.GetCollection(sr.MongoDB, "tickets")
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
	actionText := "status"
	if statusType == "change-priority" {
		actionText = "priority"
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ticket " + actionText + " changed successfully"})
	return

}
