package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"
	md "ticket-app/internal/app/models"
	database "ticket-app/internal/app/utils/database"
	jwt "ticket-app/internal/app/utils/jwt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		bearerToken := strings.Split(tokenString, " ")
		if len(bearerToken) == 0 || len(bearerToken) != 2 {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if (strings.ToLower(bearerToken[0])) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		accessToken := bearerToken[1]
		if accessToken == "" {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		claims, err := jwt.ValidateToken(accessToken)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		var dbResponse md.UserRetrieve
		collections := database.GetCollection(database.MongoDB, "users")
		//fmt.Println(collections.Find(c.TODO(), bson.M{}))

		err = collections.FindOne(context.TODO(), bson.M{"userName": claims.Username}).Decode(&dbResponse)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
		user := map[string]interface{}{"username": dbResponse.UserName, "timezone": dbResponse.TimeZone}
		c.Set("user", user)
		c.Next()
	}
}
