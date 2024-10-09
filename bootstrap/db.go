package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ETjoel/task_managment_api/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(env Env) *mongo.Client {
	cxt, cancel := context.WithTimeout(context.Background(), time.Duration(env.ContextTimeout)*time.Second)
	defer cancel()

	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%s", env.DBUser, env.DBPass, env.DBHost, env.DBPort)

	if env.DBUser == "" || env.DBPass != "" {
		mongoUri = fmt.Sprintf("mongodb://%s:%s", env.DBHost, env.DBPort)

	}

	clientOptions := options.Client().ApplyURI(mongoUri)
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

func CreateEmailUniqueIndex(env Env, client *mongo.Client) {
	collection := client.Database(env.DBName).Collection(domain.UsersCollection)

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
