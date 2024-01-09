package usecases

import (
	"errors"
	"fiber_simple_product_management_api/helpers"
	"fiber_simple_product_management_api/models"
	"fiber_simple_product_management_api/repositories"
)

type (
	UserUsecase interface {
		Register(requestRegister *models.RequestRegisterUser) (*models.User, error)
		Login(requestLogin *models.RequestLoginUser) (*helpers.ResponseJWT, error)
		GetAll() ([]models.User, error)
	}

	userUsecase struct {
		UserRepository repositories.UserRepository
	}
)

func NewUserUsecase(ur repositories.UserRepository) UserUsecase {
	return &userUsecase{
		UserRepository: ur,
	}
}

func (uu userUsecase) Register(requestRegister *models.RequestRegisterUser) (*models.User, error) {
	// hash password
	hashedPassword, err := helpers.NewHashBcryptHelper(requestRegister.Password, "").HashPassword()
	if err != nil {
		return nil, err
	}

	var user models.User
	user.Name = requestRegister.Name
	user.Email = requestRegister.Email
	user.Password = hashedPassword

	return uu.UserRepository.Create(&user)
}

func (uu userUsecase) Login(requestLogin *models.RequestLoginUser) (*helpers.ResponseJWT, error) {
	var responseJWT helpers.ResponseJWT

	user, err := uu.UserRepository.GetByEmail(requestLogin.Email)
	if err != nil {
		return nil, err
	}

	if user.Email == "" {
		return nil, errors.New("wrong email")
	}

	// hash password
	match := helpers.NewHashBcryptHelper(requestLogin.Password, user.Password).CheckPasswordHash()
	if match != nil {
		return nil, match
	}

	// generate jwt token
	responseJWT, err = helpers.NewJWTHelper(user).GenerateJWTToken()
	if err != nil {
		return nil, err
	}

	return &responseJWT, nil
}

func (uu userUsecase) GetAll() ([]models.User, error) {
	return uu.UserRepository.GetAll()
}
