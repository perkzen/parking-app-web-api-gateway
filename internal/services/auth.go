package services

import (
	"fmt"
	"parking-app-web-api-gateway/internal/auth"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Login(username, password string) (string, error) {
	if username != "admin" || password != "admin" {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := auth.GenerateJwtToken(username)

	return token, err
}
