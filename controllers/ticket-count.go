package controllers

import (
	"context"
	"net/http"
	sr "ticket-app/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TicketCount(c *gin.Context) {

	ticketCollections := sr.GetCollection(sr.MongoDB, "tickets")
	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$status"},
			{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}}
	cursor, err := ticketCollections.Aggregate(context.TODO(), mongo.Pipeline{groupStage})
	if err != nil {
		panic(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	response := map[string]interface{}{}
	for _, result := range results {
		response[result["_id"].(string)] = result["count"]
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
	return
}
