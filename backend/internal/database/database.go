package database

import (
	"context"
	"os"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() {
	username := url.QueryEscape(os.Getenv("DATABASE_USERNAME"))
	password := url.QueryEscape(os.Getenv("DATABASE_PASSWORD"))
	cluster := os.Getenv("DATABASE_CLUSTER")
	query_parameters := "retryWrites=true&w=majority"
	cluster_name := os.Getenv("DATABASE_CLUSTER_NAME")

	url := "mongodb+srv://" + username + ":" + password + "@" + cluster + "/?" + query_parameters + "&appName=" + cluster_name
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	
	Ping()
}

func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
