package main

import (
	"os"

	"github.com/mesxx/Fiber_Simple_Product_Management_API/servers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	port := os.Getenv("PORT")

	server := servers.Server()
	server.Listen(":" + port)
}
