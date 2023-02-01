package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	cf "ticket-app/commonfunctions"
	ct "ticket-app/constants"
	md "ticket-app/models"
	sr "ticket-app/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListTickets(c *gin.Context) {

	var validate = validator.New()

	var TicketListRequest md.TicketListRequest
	if err := c.ShouldBindJSON(&TicketListRequest); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}
	if err := validate.Struct(TicketListRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorResponse := cf.TranslateError(validationErrors, validate)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorResponse})
		return
	}
	/*val := reflect.ValueOf(TicketListRequest)
	field, isField := val.Type().FieldByName(fieldName)
	if isField {
		fieldJSONName = field.Tag.Get("json")
	}*/

	filterCondition := bson.D{}
	if TicketListRequest.Status != "" {
		//filter := bson.D{{"type", bson.D{{"$eq", "Oolong"}}}}
		filterCondition = append(filterCondition, bson.E{Key: "status", Value: bson.D{{Key: "$eq", Value: TicketListRequest.Status}}})
	}
	if TicketListRequest.Priority != "" {
		//filter := bson.D{{"type", bson.D{{"$eq", "Oolong"}}}}
		filterCondition = append(filterCondition, bson.E{Key: "priority", Value: bson.D{{Key: "$eq", Value: TicketListRequest.Priority}}})
	}
	if TicketListRequest.Search != "" {

		likeSearch := TicketListRequest.Search

		/*pipeline := bson.D{
			{"key1", 1},
			{"$or", []interface{}{
				bson.D{{"key2", 2}},
				bson.D{{"key3", 2}},
			}},
		}*/
		// setElements = append(setElements, bson.E{"email", pivot.Email})
		filterCondition = append(filterCondition, bson.E{Key: "$or", Value: bson.A{
			//bson.D{{Key: "description", Value: bson.M{"$regex": likeSearch, "$options": "im"}}},
			bson.D{{Key: "subject", Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: likeSearch, Options: "i"}}}}},
			bson.D{{Key: "description", Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: likeSearch, Options: "i"}}}}},
			bson.D{{Key: "contact.name", Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: likeSearch, Options: "i"}}}}},
			bson.D{{Key: "contact.lastname", Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: likeSearch, Options: "i"}}}}},
			bson.D{{Key: "contact.email", Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: likeSearch, Options: "i"}}}}},
			bson.D{{Key: "assignee.name", Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: likeSearch, Options: "i"}}}}},
			bson.D{{Key: "assignee.lastname", Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: likeSearch, Options: "i"}}}}},
			bson.D{{Key: "assignee.email", Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: likeSearch, Options: "i"}}}}},
		}})

	}
	userDt, _ := c.Get("user")
	user := userDt.(map[string]interface{})
	if TicketListRequest.DateFrom != "" && TicketListRequest.DateTo != "" {
		dateFrom := cf.ConvertToUtcDateTime(TicketListRequest.DateFrom+"T00:00:00.000Z", user["timezone"].(string), "2006-01-02T15:04:05.000Z")
		dateTo := cf.ConvertToUtcDateTime(TicketListRequest.DateTo+"T23:59:59.999Z", user["timezone"].(string), "2006-01-02T15:04:05.000Z")
		t1, _ := time.Parse(ct.ISO_DB_DATE_FORMAT, dateFrom) //converted to ISODate format
		t2, _ := time.Parse(ct.ISO_DB_DATE_FORMAT, dateTo)   //converted to ISODate format

		filterCondition = append(filterCondition, bson.E{Key: "$and", Value: bson.A{
			bson.D{{Key: "createdDate", Value: bson.D{{Key: "$gte", Value: t1}}}},
			bson.D{{Key: "createdDate", Value: bson.D{{Key: "$lte", Value: t2}}}},
		}})
	}
	//fmt.Println(filterCondition)

	ticketCollections := sr.GetCollection(sr.MongoDB, "tickets")
	ticketCount, err := ticketCollections.CountDocuments(context.TODO(), filterCondition)

	opts := options.Find()
	sortField := "createdDate"
	sortOrder := -1
	if TicketListRequest.SortField != "" && TicketListRequest.SortOrder != "" {
		sortField = TicketListRequest.SortField
		if TicketListRequest.SortOrder == "asc" {
			sortOrder = 1
		}
	}
	opts.SetSort(bson.D{{Key: sortField, Value: sortOrder}})

	offset := (TicketListRequest.Page - 1) * TicketListRequest.Limit
	opts.SetSkip(int64(offset))
	opts.SetLimit(int64(TicketListRequest.Limit))

	results, err := ticketCollections.Find(context.TODO(), filterCondition, opts)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	defer results.Close(context.TODO())
	var tickets []md.TicketJson

	for results.Next(context.TODO()) {
		var singleTicket md.TicketJson
		if err = results.Decode(&singleTicket); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		//fmt.Println(singleTicket)
		CreatedDate := cf.ConvertUtcDateTime(fmt.Sprint(singleTicket.CreatedDate.Time().UTC().Format(ct.API_DATE_RESPONSE_FORMAT)), user["timezone"].(string), ct.API_DATE_RESPONSE_FORMAT)
		singleTicket.DateCreated = CreatedDate
		tickets = append(tickets, singleTicket)

	}

	c.JSON(http.StatusOK, gin.H{"count": ticketCount, "listCount": len(tickets), "data": tickets})
	return
}
