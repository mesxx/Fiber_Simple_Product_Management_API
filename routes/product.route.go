package routes

import (
	"github.com/mesxx/Fiber_Simple_Product_Management_API/handlers"
	"github.com/mesxx/Fiber_Simple_Product_Management_API/middlewares"
	"github.com/mesxx/Fiber_Simple_Product_Management_API/repositories"
	"github.com/mesxx/Fiber_Simple_Product_Management_API/usecases"

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
