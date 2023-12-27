package models

type (
	RequestRegisterUser struct {
		Name     string `json:"name" validate:"required,min=5,max=20"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=5"`
	}

	RequestLoginUser struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=5"`
	}

	RequestCreateProduct struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		Price       uint   `json:"price" validate:"required"`
		Stock       uint   `json:"stock" validate:"required"`
	}

	RequestUpdateProduct struct {
		Name        *string `json:"name"`
		Description *string `json:"description"`
		Price       *uint   `json:"price"`
		Stock       *uint   `json:"stock"`
	}
)
