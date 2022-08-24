package main

import (
	"github.com/milisecond/ecommerce-service/tree/main/user/app"
     "github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
)

func main(){
     godotenv.Load(".env")
     _ = app.NewMongo().NewDB().(*mongo.Client)
}