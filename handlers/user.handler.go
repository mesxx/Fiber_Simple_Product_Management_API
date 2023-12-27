package handlers

import (
	"fiber_simple_product_management/helpers"
	"fiber_simple_product_management/models"
	"fiber_simple_product_management/usecases"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	UserHandler interface {
		Register(c *fiber.Ctx) error
		Login(c *fiber.Ctx) error
		GetAllUser(c *fiber.Ctx) error
	}

	userHandler struct {
		UserUsecase usecases.UserUsecase
	}
)

func NewUserHandler(uu usecases.UserUsecase) UserHandler {
	return &userHandler{
		UserUsecase: uu,
	}
}

func (uh userHandler) Register(c *fiber.Ctx) error {
	var requestRegister models.RequestRegisterUser
	if err := c.BodyParser(&requestRegister); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validator
	validate := validator.New()
	if err := validate.Struct(requestRegister); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// end validator

	res, err := uh.UserUsecase.Register(&requestRegister)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return helpers.NewResponseDataHelper(fiber.StatusCreated, "success created", res).GetResponse(c)
}

func (uh userHandler) Login(c *fiber.Ctx) error {
	var requestLogin models.RequestLoginUser
	if err := c.BodyParser(&requestLogin); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validator
	validate := validator.New()
	if err := validate.Struct(requestLogin); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// end validator

	res, err := uh.UserUsecase.Login(&requestLogin)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return helpers.NewResponseDataHelper(fiber.StatusOK, "success login", res).GetResponse(c)
}

func (uh userHandler) GetAllUser(c *fiber.Ctx) error {
	users, err := uh.UserUsecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return helpers.NewResponseDataHelper(fiber.StatusOK, "success", users).GetResponse(c)
}
