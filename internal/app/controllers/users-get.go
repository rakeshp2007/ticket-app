package controllers

import (
	"context"
	"log"
	"net/http"
	md "ticket-app/internal/app/models"
	database "ticket-app/internal/app/utils/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers(c *gin.Context) {
	userCollections := database.GetCollection(database.MongoDB, "users")
	opts := options.Find().SetProjection(bson.D{{Key: "name", Value: 1}, {Key: "userName", Value: 1}})
	results, err := userCollections.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"eror": err.Error()})
		return
	}
	defer results.Close(context.TODO())
	var users []md.UserRetrieve
	for results.Next(context.TODO()) {
		var singleUser md.UserRetrieve
		if err = results.Decode(&singleUser); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		users = append(users, singleUser)
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
	return
}
