package usecases

import (
	"errors"
	"fiber_simple_product_management/models"
	"fiber_simple_product_management/repositories"
)

type (
	ProductUsecase interface {
		Create(product *models.RequestCreateProduct, user *models.CustomClaims) (*models.Product, error)
		GetAll() ([]models.Product, error)
		GetByID(id uint) (*models.Product, error)
		Update(requestProduct *models.RequestUpdateProduct, id uint) (*models.Product, error)
		Delete(id uint) (*models.Product, error)
	}

	productUsecase struct {
		ProductRepository repositories.ProductRepository
	}
)

func NewProductUsecase(pr repositories.ProductRepository) ProductUsecase {
	return &productUsecase{
		ProductRepository: pr,
	}
}

func (pu productUsecase) Create(requestCreateProduct *models.RequestCreateProduct, user *models.CustomClaims) (*models.Product, error) {
	var product models.Product

	product.Name = requestCreateProduct.Name
	product.Description = requestCreateProduct.Description
	product.Price = requestCreateProduct.Price
	product.Stock = requestCreateProduct.Stock
	product.UserId = user.ID

	return pu.ProductRepository.Create(&product)
}

func (pu productUsecase) GetAll() ([]models.Product, error) {
	return pu.ProductRepository.GetAll()
}

func (pu productUsecase) GetByID(id uint) (*models.Product, error) {
	return pu.ProductRepository.GetByID(id)
}

func (pu productUsecase) Update(requestUpdateProduct *models.RequestUpdateProduct, id uint) (*models.Product, error) {
	getProduct, err := pu.ProductRepository.GetByID(id)
	if err != nil {
		return nil, err
	} else if getProduct.ID == 0 {
		return nil, errors.New("product ID is invalid, please try again")
	}

	if requestUpdateProduct.Name != nil {
		getProduct.Name = *requestUpdateProduct.Name
	}

	if requestUpdateProduct.Description != nil {
		getProduct.Description = *requestUpdateProduct.Description
	}

	if requestUpdateProduct.Price != nil {
		getProduct.Price = *requestUpdateProduct.Price
	}

	if requestUpdateProduct.Stock != nil {
		getProduct.Stock = *requestUpdateProduct.Stock
	}

	updateProduct, err := pu.ProductRepository.Update(getProduct)
	if err != nil {
		return nil, err
	}

	return updateProduct, nil
}

func (pu productUsecase) Delete(id uint) (*models.Product, error) {
	getProduct, err := pu.ProductRepository.GetByID(id)
	if err != nil {
		return nil, err
	} else if getProduct.ID == 0 {
		return nil, errors.New("product ID is invalid, please try again")
	}

	deleteProduct, err := pu.ProductRepository.Delete(getProduct, getProduct.ID)
	if err != nil {
		return nil, err
	}

	return deleteProduct, nil
}
