package app

import (
	"context"
	"fmt"
	"github.com/milisecond/ecommerce-service/helper"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
}

func NewMongo() database {
	return &mongoDB{}
}

func (md *mongoDB) NewDB() interface{} {
	fmt.Println("Connecting Database...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	_ = cancel

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		Username:      os.Getenv("DB_USERNAME"), // your mongodb user
		Password:      os.Getenv("DB_PASSWORD"), // ...and mongodb
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_HOST")).SetAuth(credential)

	client, err := mongo.NewClient(clientOptions)
	helper.PanicIfError(err)
	// connect to mongo
	err = client.Connect(ctx)
	helper.PanicIfError(err)
	// test connection to mongo
	err = client.Ping(ctx, nil)
	helper.PanicIfError(err)
	fmt.Println("Connected")

	return client
}
