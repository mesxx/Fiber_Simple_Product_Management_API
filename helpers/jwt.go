package helpers

import (
	"fiber_simple_product_management/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	JWTHelper interface {
		GenerateJWTToken() (ResponseJWT, error)
	}

	JWTHelperStruct struct {
		User models.User
	}

	ResponseJWT struct {
		Token string      `json:"token"`
		User  models.User `json:"user"`
	}
)

func NewJWTHelper(user *models.User) JWTHelper {
	return &JWTHelperStruct{
		User: *user,
	}
}

func (jhs JWTHelperStruct) GenerateJWTToken() (ResponseJWT, error) {
	var responseJWT ResponseJWT
	secretKey := os.Getenv("SECRET_KEY")

	// Create the Claims
	claims := models.CustomClaims{
		ID:    jhs.User.ID,
		Name:  jhs.User.Name,
		Email: jhs.User.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return responseJWT, err
	}

	responseJWT.Token = t
	responseJWT.User = jhs.User
	return responseJWT, nil
}
