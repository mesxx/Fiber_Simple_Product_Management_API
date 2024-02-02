package routes

import (
	"github.com/mesxx/Fiber_Simple_Product_Management_API/handlers"
	"github.com/mesxx/Fiber_Simple_Product_Management_API/middlewares"
	"github.com/mesxx/Fiber_Simple_Product_Management_API/repositories"
	"github.com/mesxx/Fiber_Simple_Product_Management_API/usecases"

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
