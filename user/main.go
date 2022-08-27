package main

import (
	"github.com/milisecond/ecommerce-service/app"
     "github.com/joho/godotenv"
     "log"

     "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	userSvc "github.com/milisecond/ecommerce-service/module/user/app"
)



func main(){
     godotenv.Load(".env")
     _ = app.NewMongo().NewDB().(*mongo.Client)

     app := fiber.New()
     userSvc.Register(app)

     log.Fatal(app.Listen(":3000"))
}