package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	cf "ticket-app/commonfunctions"
	md "ticket-app/models"
	sr "ticket-app/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var inputData md.LoginInput
	var dbResponse md.LoginInput
	var validate = validator.New()
	if err := c.ShouldBindJSON(&inputData); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if validationErr := validate.Struct(&inputData); validationErr != nil {
		//c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		validationErrors := validationErr.(validator.ValidationErrors)
		errorResponse := cf.TranslateError(validationErrors, validate)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorResponse})
		return
	}

	collections := sr.GetCollection(sr.MongoDB, "users")
	//fmt.Println(collections.Find(context.TODO(), bson.M{}))

	err := collections.FindOne(context.TODO(), bson.M{"userName": inputData.Username}).Decode(&dbResponse)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	passwordIsValid, msg := VerifyPassword(inputData.Password, dbResponse.Password)
	if passwordIsValid != true {
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		return
	}

	token, err := sr.GenerateJWT(dbResponse.Username)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token creation error occured"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": token})
	return

}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("Invalid username or password")
		check = false
	}

	return check, msg
}
