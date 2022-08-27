package app

import (
	_ "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	service = "user"
	version = "v1"
	base_url = "/api/"+ service +"/" +version + "/"
)

func greeting(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("User Service here")
}

func Register(route *fiber.App) {
	route.Get(base_url + ":user_id", greeting)
	route.Post(base_url + "register", greeting)
	route.Patch(base_url + ":user_id", greeting)
}