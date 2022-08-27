package app

import (
	_ "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	service = "auth"
	version = "v1"
	base_url = "/api/"+ service +"/" +version + "/"
)

func greeting(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("Auth Service here")
}

func Register(route *fiber.App) {
	route.Post(base_url + "login", greeting)
}