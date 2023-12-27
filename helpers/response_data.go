package helpers

import (
	"fiber_simple_product_management/models"

	"github.com/gofiber/fiber/v2"
)

type (
	ResponseDataHelper interface {
		GetResponse(c *fiber.Ctx) error
	}

	responseDataHelper struct {
		ResponseData models.ResponseData
	}
)

func NewResponseDataHelper(code int, message string, data interface{}) ResponseDataHelper {
	return &responseDataHelper{
		ResponseData: models.ResponseData{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}
}

func (r responseDataHelper) GetResponse(c *fiber.Ctx) error {
	return c.Status(r.ResponseData.Code).JSON(r.ResponseData)
}
