package servers

import (
	"fiber_simple_product_management/configs"
	"fiber_simple_product_management/middlewares"
	"fiber_simple_product_management/models"
	"fiber_simple_product_management/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Server() *fiber.App {
	db, err := configs.DatabaseConfig()
	if err != nil {
		fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	db.AutoMigrate(&models.User{}, &models.Product{})

	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorMiddleware,
	})
	app.Use(logger.New())
	app.Use(recover.New())

	api := app.Group("/api")
	users := api.Group("/users")
	products := api.Group("/products")

	routes.UserRouter(users, db)
	routes.ProductRouter(products, db)

	return app
}
