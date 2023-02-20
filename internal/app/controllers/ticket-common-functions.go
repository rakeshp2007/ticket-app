package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	md "ticket-app/internal/app/models"
	cf "ticket-app/internal/app/utils/commonfunctions"
	database "ticket-app/internal/app/utils/database"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}

func ValidateId(c *gin.Context) (status bool, message string, docID primitive.ObjectID, singleTicket md.TicketJson) {

	docID, err := primitive.ObjectIDFromHex(strings.TrimSpace(c.Param("id")))

	if err != nil {
		fmt.Println(err)
		status = false
		message = err.Error()
		return
		//c.JSON(http.StatusBadRequest, gin.H{"eror": map[string]interface{}{"id": err.Error()}})
		//return
	}
	ticketCollections := database.GetCollection(database.MongoDB, "tickets")
	err = ticketCollections.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&singleTicket)
	if err != nil {
		log.Println(err)
		status = false
		message = "record not found with the given id"
		return
		//c.JSON(http.StatusBadRequest, gin.H{"eror": map[string]interface{}{"id": "record not found with the given id"}})
		//return
	}
	status = true
	return
}

func GetTicketSaveData(isEdit bool, TicketCreateValidate md.CreateTicketValidate, ContactDetails md.ContactRetrieve, UserDetails md.UserRetrieve) md.TicketJson {
	ticketSaveData := md.TicketJson{
		Subject:     TicketCreateValidate.Subject,
		Description: TicketCreateValidate.Description,
		Status:      TicketCreateValidate.Status,
		Priority:    TicketCreateValidate.Priority,

		Contact: Contact{
			Id:       fmt.Sprint(TicketCreateValidate.Contact),
			Name:     ContactDetails.Name,
			LastName: ContactDetails.LastName,
			Email:    ContactDetails.Email,
		},
		Assignee: Assignee{
			Id:       fmt.Sprint(TicketCreateValidate.Assignee),
			Name:     UserDetails.Name,
			LastName: UserDetails.LastName,
			Email:    UserDetails.UserName,
		},
	}
	if !isEdit {
		dt, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		ticketSaveData.CreatedDate = primitive.NewDateTimeFromTime(dt)
	}
	return ticketSaveData
}

func ValidateTicket(c *gin.Context) (status bool, TicketCreateValidate md.CreateTicketValidate) {
	var validate = validator.New()

	if err := c.ShouldBindJSON(&TicketCreateValidate); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		status = false
		return
	}

	if validationErr := validate.Struct(&TicketCreateValidate); validationErr != nil {
		//c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		validationErrors := validationErr.(validator.ValidationErrors)
		errorResponse := cf.TranslateError(validationErrors, validate)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorResponse})
		status = false
		return
	}
	TicketCreateValidate.Contact = strings.TrimSpace(fmt.Sprint(TicketCreateValidate.Contact))
	TicketCreateValidate.Assignee = strings.TrimSpace(fmt.Sprint(TicketCreateValidate.Assignee))
	TicketCreateValidate.Subject = strings.TrimSpace(fmt.Sprint(TicketCreateValidate.Subject))
	TicketCreateValidate.Description = strings.TrimSpace(fmt.Sprint(TicketCreateValidate.Description))
	status = true
	return
}

func ValidateContact(contact string) (status bool, message string, ContactDetails md.ContactRetrieve) {

	contactID, err := primitive.ObjectIDFromHex(contact)

	if err != nil {
		log.Println(err)
		//c.JSON(http.StatusBadRequest, gin.H{"eror": map[string]interface{}{"Contact": err.Error()}})
		status = false
		message = err.Error()
		return
	}
	contactCollections := database.GetCollection(database.MongoDB, "contacts")
	err = contactCollections.FindOne(context.TODO(), bson.M{"_id": contactID}).Decode(&ContactDetails)
	if err != nil {
		log.Println(err)
		status = false
		message = "record not found with the given id"
		//c.JSON(http.StatusBadRequest, gin.H{"eror": map[string]interface{}{"Contact": "record not found with the given id"}})
		return
	}
	status = true
	return

}
func validateUser(assignee string) (status bool, message string, UserDetails md.UserRetrieve) {

	userId, err := primitive.ObjectIDFromHex(assignee)

	if err != nil {
		log.Println(err)
		status = false
		message = err.Error()
		return
		//c.JSON(http.StatusBadRequest, gin.H{"eror": map[string]interface{}{"Assignee": err.Error()}})
	}

	userCollections := database.GetCollection(database.MongoDB, "users")
	err = userCollections.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(&UserDetails)
	if err != nil {
		log.Println(err)
		status = false
		message = "record not found with the given id"
		//c.JSON(http.StatusBadRequest, gin.H{"eror": map[string]interface{}{"Assignee": "record not found with the given id"}})
		return

	}
	status = true
	return
}

func ValidateTicketChange(c *gin.Context) (status bool, TicketChangeValidate md.TicketChange) {
	var validate = validator.New()

	if err := c.ShouldBindJSON(&TicketChangeValidate); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		status = false
		return
	}

	if validationErr := validate.Struct(&TicketChangeValidate); validationErr != nil {
		//c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		validationErrors := validationErr.(validator.ValidationErrors)
		errorResponse := cf.TranslateError(validationErrors, validate)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorResponse})
		status = false
		return
	}
	status = true
	return
}

func GetTicketChangeData(changeType string, TicketChange md.TicketChange) md.TicketChange {
	ticketChangeData := md.TicketChange{}
	if changeType == "change-status" && TicketChange.Status != "" {
		ticketChangeData.Status = TicketChange.Status
	} else if changeType == "change-priority" && TicketChange.Priority != "" {
		ticketChangeData.Priority = TicketChange.Priority
	}
	return ticketChangeData
}
