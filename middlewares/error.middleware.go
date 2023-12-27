package middlewares

import (
	"errors"
	"fiber_simple_product_management/helpers"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ErrorMiddleware(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	} else if strings.Contains(err.Error(), "bad") {
		code = 400
		message = err.Error()
	}

	return helpers.NewResponseErrorHelper(code, message).GetResponse(c)
}
