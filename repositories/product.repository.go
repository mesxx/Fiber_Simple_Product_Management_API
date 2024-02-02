package repositories

import (
	"github.com/mesxx/Fiber_Simple_Product_Management_API/models"

	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		Create(product *models.Product) (*models.Product, error)
		GetAll() ([]models.Product, error)
		GetByID(id uint) (*models.Product, error)
		Update(product *models.Product) (*models.Product, error)
		Delete(product *models.Product, id uint) (*models.Product, error)
	}

	productRepository struct {
		DB *gorm.DB
	}
)

func NewProductRepositoy(db *gorm.DB) ProductRepository {
	return &productRepository{
		DB: db,
	}
}

func (pr productRepository) Create(product *models.Product) (*models.Product, error) {
	if err := pr.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (pr productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := pr.DB.Preload("User").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pr productRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := pr.DB.Preload("User").Where("ID = ?", id).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (pr productRepository) Update(product *models.Product) (*models.Product, error) {
	if err := pr.DB.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (pr productRepository) Delete(product *models.Product, id uint) (*models.Product, error) {
	if err := pr.DB.Where("ID = ?", id).Delete(product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}
