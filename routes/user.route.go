package routes

import (
	"fiber_simple_product_management_api/handlers"
	"fiber_simple_product_management_api/middlewares"
	"fiber_simple_product_management_api/repositories"
	"fiber_simple_product_management_api/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRouter(user fiber.Router, db *gorm.DB) {
	ur := repositories.NewUserRepositoy(db)
	uu := usecases.NewUserUsecase(ur)
	uh := handlers.NewUserHandler(uu)

	user.Post("/register", uh.Register)
	user.Post("/Login", uh.Login)

	// Authorization
	user.Use(middlewares.RestrictedUser)
	user.Get("/", uh.GetAllUser)
}
