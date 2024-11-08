package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Message struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Content  string             `json:"content" bson:"content"`
}

func GetMessages(channel string, amount int64) []Message {
	collection := client.Database("channels").Collection(channel)
	cursor, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetLimit(amount))
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.TODO())

	var documents []Message
	for cursor.Next(context.TODO()) {
		var document Message
		if err := cursor.Decode(&document); err != nil {
			log.Fatal(err)
		}
		documents = append(documents, document)
	}

	return documents
}

func SendMessage(channel string, message Message) {
	collection := client.Database("channels").Collection(channel)
	collection.InsertOne(context.TODO(), bson.D{{Key: "username", Value: message.Username}, {Key: "content", Value: message.Content}})
}

func WatchChannel(channel string) *mongo.ChangeStream {
	collection := client.Database("channels").Collection(channel)
	pipeline := mongo.Pipeline{bson.D{{Key: "$match", Value: bson.D{{Key: "operationType", Value: bson.D{{Key: "$in", Value: bson.A{"insert", "delete"}}}}}}}}

	cs, err := collection.Watch(context.TODO(), pipeline)
	if err != nil {
		log.Fatal("Database error: ", err)
	}
  
	return cs
}
