package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func RunCommand(runCommand interface{}) {
	if err := client.Database("admin").RunCommand(context.TODO(), runCommand).Err(); err != nil {
    panic(err)
  }
}

func Ping() {
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
    panic(err)
  }
  fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
