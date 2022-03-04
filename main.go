package main

import (
	"github.com/cecep31/fiberapi/config"
	"github.com/cecep31/fiberapi/exception"
	"github.com/cecep31/fiberapi/migrations"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//setup configurasi
	cofiguration := config.New()
	database := config.NewPgsql(cofiguration)

	//auto migrations
	migrations.ProduckMigration(database)

	//setup controller
	// produckcontroller := service.NewPro

	//setup fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Jakarta",
	}))

	err := app.Listen(":1234")
	exception.PanicIfNeeded(err)
}
