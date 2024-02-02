package repositories

import (
	"github.com/mesxx/Fiber_Simple_Product_Management_API/models"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(user *models.User) (*models.User, error)
		GetByEmail(email string) (*models.User, error)
		GetAll() ([]models.User, error)
	}

	userRepository struct {
		DB *gorm.DB
	}
)

func NewUserRepositoy(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur userRepository) Create(user *models.User) (*models.User, error) {
	if err := ur.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.DB.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := ur.DB.Preload("Products").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
