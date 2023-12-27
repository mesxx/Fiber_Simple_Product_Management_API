package helpers

import "golang.org/x/crypto/bcrypt"

type (
	HashBcryptHelper interface {
		HashPassword() (string, error)
		CheckPasswordHash() error
	}

	HashBcryptHelperStruct struct {
		Password       string
		HashedPassword string
	}
)

func NewHashBcryptHelper(password string, hashedPassword string) HashBcryptHelper {
	return &HashBcryptHelperStruct{
		Password:       password,
		HashedPassword: hashedPassword,
	}
}

func (hbs HashBcryptHelperStruct) HashPassword() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(hbs.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (hbs HashBcryptHelperStruct) CheckPasswordHash() error {
	err := bcrypt.CompareHashAndPassword([]byte(hbs.HashedPassword), []byte(hbs.Password))
	return err
}
