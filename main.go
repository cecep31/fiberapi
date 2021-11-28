package main

import (
	"github.com/cecep31/fiberapi/config"
	"github.com/cecep31/fiberapi/exception"
	"github.com/cecep31/fiberapi/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//setup configurasi
	cofiguration := config.New()
	database := config.NewMysql(cofiguration)

	//setup controller
	produckcontroller := service.NewPro

	//setup fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	err := app.Listen(":1234")
	exception.PanicIfNeeded(err)
}
