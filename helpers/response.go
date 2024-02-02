package helpers

import (
	"github.com/mesxx/Fiber_Simple_Product_Management_API/models"

	"github.com/gofiber/fiber/v2"
)

type (
	ResponseErrorHelper interface {
		GetResponse(c *fiber.Ctx) error
	}

	responseErrorHelper struct {
		Response models.Response
	}
)

func NewResponseErrorHelper(code int, message string) ResponseErrorHelper {
	return &responseErrorHelper{
		Response: models.Response{
			Code:    code,
			Message: message,
		},
	}
}

func (r responseErrorHelper) GetResponse(c *fiber.Ctx) error {
	return c.Status(r.Response.Code).JSON(r.Response)
}
