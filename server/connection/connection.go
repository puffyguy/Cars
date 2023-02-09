package connection

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDB() *mongo.Database {
	credential := options.Credential{
		AuthSource: "admin",
		Username:   "mongoadmin",
		Password:   "mongoadmin",
	}
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/").SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("Praful")
}
