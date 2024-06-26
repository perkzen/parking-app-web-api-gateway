package routes

import (
	"github.com/gin-gonic/gin"
	"parking-app-web-api-gateway/internal/auth"
	"parking-app-web-api-gateway/internal/handlers"
)

func setupParkingRoutes(app *gin.Engine) {

	parkingHandler := handlers.NewParkingHandler()

	app.Use(auth.AuthMiddleware())
	app.GET("/park", parkingHandler.GetParkingSpots)
	app.GET("/park/:name", parkingHandler.GetParkingSpot)
	app.POST("/park", parkingHandler.CreateParkingSpot)
}
