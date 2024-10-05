package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	TaskCollection  = "tasks"
	UsersCollection = "users"
)

var client *mongo.Client

func ConnectMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	return client.Database("task_manager").Collection(collectionName)
}
func GetAuthCollection(collectionName string) *mongo.Collection {
	return client.Database("task_manager").Collection(collectionName)
}

func CreateEmailUniqueIndex() {
	collection := GetCollection(UsersCollection)

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatalf("Coudn't make email unique: %s", err.Error())
	}
}
func DisconnectMongoDB() {
	err := client.Disconnect(context.Background())

	if err != nil {
		log.Fatal(err)
	}
}
