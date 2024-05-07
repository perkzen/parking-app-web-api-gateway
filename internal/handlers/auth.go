package handlers

import (
	"github.com/gin-gonic/gin"
	"parking-app-web-api-gateway/internal/services"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		AuthService: services.NewAuthService(),
	}
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login godoc
// @Tags auth
// @Summary Login
// @Description Login
// @Accept json
// @Produce json
// @Param loginDTO body LoginDTO true "Login DTO"
// @Success 200 {string} string "JWT token"
// @Router /login [post]
func (a *AuthHandler) Login(c *gin.Context) {
	var loginDTO LoginDTO
	err := c.BindJSON(&loginDTO)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := a.AuthService.Login(loginDTO.Username, loginDTO.Password)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
