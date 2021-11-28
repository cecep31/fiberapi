package exception

import (
	"github.com/cecep31/fiberapi/model"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return c.JSON(model.WebResponse{
			Code:   400,
			Status: "bad request",
			Data:   err.Error(),
		})
	}
	return c.JSON(model.WebResponse{
		Code:   500,
		Status: "internal server error",
		Data:   err.Error(),
	})
}
