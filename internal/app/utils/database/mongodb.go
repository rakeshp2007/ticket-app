package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client
var MongoDBName string

func ConnectDB(MongoDBStruct struct {
	Host     string
	Username string
	Password string
	Database string
	Port     string
}) *mongo.Client {

	MongoDBName = url.QueryEscape(MongoDBStruct.Database)
	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", url.QueryEscape(MongoDBStruct.Username), url.QueryEscape(MongoDBStruct.Password), url.QueryEscape(MongoDBStruct.Host), url.QueryEscape(MongoDBStruct.Port), MongoDBName)
	//mongoUri := "mongodb://ticketUser:User%40123@localhost:27017/ticket-app"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	//dbb, _ := client.ListDatabases(ctx, bson.M{})
	//fmt.Println(dbb)
	//fmt.Println("Connected to MongoDB")
	MongoDB = client
	return client
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(MongoDBName).Collection(collectionName)
	return collection
}
