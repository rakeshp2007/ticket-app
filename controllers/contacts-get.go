package controllers

import (
	"context"
	"log"
	"net/http"
	md "ticket-app/models"
	sr "ticket-app/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetContacts(c *gin.Context) {
	contactCollections := sr.GetCollection(sr.MongoDB, "contacts")
	opts := options.Find().SetProjection(bson.D{{Key: "name", Value: 1}, {Key: "email", Value: 1}})
	results, err := contactCollections.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"eror": err.Error()})
		return
	}
	defer results.Close(context.TODO())
	var contacts []md.ContactRetrieve
	for results.Next(context.TODO()) {
		var singleContact md.ContactRetrieve
		if err = results.Decode(&singleContact); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		contacts = append(contacts, singleContact)
	}
	c.JSON(http.StatusOK, gin.H{"data": contacts})
	return
}
