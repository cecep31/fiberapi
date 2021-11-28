package controller

import (
	"github.com/cecep31/fiberapi/exception"
	"github.com/cecep31/fiberapi/model"
	"github.com/cecep31/fiberapi/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProduckController struct {
	ProduckService service.ProduckService
}

func NewProduckController(producksevice *service.ProduckService) ProduckController {
	return ProduckController{ProduckService: *producksevice}
}

func (controller *ProduckController) Route(app *fiber.App) {
	app.Post("/api/produck", controller.Create)
	app.Get("/api/produck", controller.List)
}

func (controller *ProduckController) Create(c *fiber.Ctx) error {
	var request model.CreateProduckRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()
	exception.PanicIfNeeded(err)

	response := controller.ProduckService.Create(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   response,
	})
}

func (controller *ProduckController) List(c *fiber.Ctx) error {
	response := controller.ProduckService.List()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   response,
	})
}
