package controllers

import (
	"fmt"
	"net/http"
	cf "ticket-app/internal/app/utils/commonfunctions"
	ct "ticket-app/internal/app/utils/constants"

	"github.com/gin-gonic/gin"
)

func GetTicket(c *gin.Context) {
	userDt, _ := c.Get("user")
	user := userDt.(map[string]interface{})
	validateStat, validateMessage, _, singleTicket := ValidateId(c)
	if !validateStat {
		c.JSON(http.StatusBadRequest, gin.H{"eror": map[string]interface{}{"Contact": validateMessage}})
		return
	}

	CreatedDate := cf.ConvertUtcDateTime(fmt.Sprint(singleTicket.CreatedDate.Time().UTC().Format(ct.API_DATE_RESPONSE_FORMAT)), user["timezone"].(string), ct.API_DATE_RESPONSE_FORMAT)
	singleTicket.DateCreated = CreatedDate
	c.JSON(http.StatusOK, gin.H{"data": singleTicket})
	return
}
