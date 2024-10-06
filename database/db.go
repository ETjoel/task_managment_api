package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	TaskCollection  = "tasks"
	UsersCollection = "users"
)

func ConnectMongoDB() *mongo.Client {
	cxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(cxt, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(cxt, nil)

	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("task_manager").Collection(collectionName)
}

func CreateEmailUniqueIndex(client *mongo.Client) {
	collection := GetCollection(client, UsersCollection)

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatalf("Coudn't make email unique: %s", err.Error())
	}
}
func DisconnectMongoDB(client *mongo.Client) {
	err := client.Disconnect(context.Background())

	if err != nil {
		log.Fatal(err)
	}
}
