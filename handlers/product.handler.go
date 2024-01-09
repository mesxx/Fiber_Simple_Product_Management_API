package handlers

import (
	"fiber_simple_product_management_api/helpers"
	"fiber_simple_product_management_api/models"
	"fiber_simple_product_management_api/usecases"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	ProductHandler interface {
		CreateProduct(c *fiber.Ctx) error
		GetAllProduct(c *fiber.Ctx) error
		GetProductByID(c *fiber.Ctx) error
		UpdateProduct(c *fiber.Ctx) error
		DeleteProduct(c *fiber.Ctx) error
	}

	productHandler struct {
		ProductUsecase usecases.ProductUsecase
	}
)

func NewProductHandler(pu usecases.ProductUsecase) ProductHandler {
	return &productHandler{
		ProductUsecase: pu,
	}
}

func (ph productHandler) CreateProduct(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.CustomClaims)

	var requestCreateProduct models.RequestCreateProduct
	if err := c.BodyParser(&requestCreateProduct); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validator
	validate := validator.New()
	if err := validate.Struct(requestCreateProduct); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// end validator

	res, err := ph.ProductUsecase.Create(&requestCreateProduct, user)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return helpers.NewResponseDataHelper(fiber.StatusCreated, "success created", res).GetResponse(c)
}

func (ph productHandler) GetAllProduct(c *fiber.Ctx) error {
	products, err := ph.ProductUsecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return helpers.NewResponseDataHelper(fiber.StatusOK, "success", products).GetResponse(c)
}

func (ph productHandler) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	value, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	product, err := ph.ProductUsecase.GetByID(uint(value))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return helpers.NewResponseDataHelper(fiber.StatusOK, "success", product).GetResponse(c)
}

func (ph productHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	value, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var requestUpdateProduct models.RequestUpdateProduct
	if err := c.BodyParser(&requestUpdateProduct); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	product, err := ph.ProductUsecase.Update(&requestUpdateProduct, uint(value))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return helpers.NewResponseDataHelper(fiber.StatusOK, "success", product).GetResponse(c)
}

func (ph productHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	value, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	product, err := ph.ProductUsecase.Delete(uint(value))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return helpers.NewResponseDataHelper(fiber.StatusOK, "success", product).GetResponse(c)
}
