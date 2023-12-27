package routes

import (
	"fiber_simple_product_management/handlers"
	"fiber_simple_product_management/middlewares"
	"fiber_simple_product_management/repositories"
	"fiber_simple_product_management/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProductRouter(product fiber.Router, db *gorm.DB) {
	pr := repositories.NewProductRepositoy(db)
	pu := usecases.NewProductUsecase(pr)
	ph := handlers.NewProductHandler(pu)

	// Authorization
	product.Use(middlewares.RestrictedUser)
	product.Post("/", ph.CreateProduct)
	product.Get("/", ph.GetAllProduct)
	product.Get("/:id", ph.GetProductByID)
	product.Put("/:id", ph.UpdateProduct)
	product.Delete("/:id", ph.DeleteProduct)
}
