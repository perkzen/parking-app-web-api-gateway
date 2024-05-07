package routes

import (
	"github.com/gin-gonic/gin"
	"parking-app-web-api-gateway/internal/handlers"
)

func setupAuthRoutes(app *gin.Engine) {

	authHandler := handlers.NewAuthHandler()

	app.POST("/login", authHandler.Login)
}
