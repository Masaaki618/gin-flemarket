package services

import (
	"gin-flemarket/models"
	"gin-flemarket/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Signup(email string, password string) error
}

type AuthService struct {
	repositories repositories.IAuthRepository
}

func NewAuthService(repositories repositories.IAuthRepository) IAuthService {
	return &AuthService{repositories: repositories}
}

func (s *AuthService) Signup(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.repositories.CreateUser(user)
}
