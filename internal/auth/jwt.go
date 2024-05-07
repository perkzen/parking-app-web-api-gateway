package auth

import "github.com/dgrijalva/jwt-go"

func GenerateJwtToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = username

	tokenString, err := token.SignedString([]byte("secret"))

	return tokenString, err
}
