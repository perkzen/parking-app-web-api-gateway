package routes

import (
	"github.com/gin-gonic/gin"
	"parking-app-web-api-gateway/internal/auth"
	"parking-app-web-api-gateway/internal/handlers"
)

func setupPaymentRoutes(app *gin.Engine) {

	paymentHandler := handlers.NewPaymentHandler()

	app.Use(auth.AuthMiddleware())
	app.GET("/payments", paymentHandler.GetPayments)
	app.PUT("/payments", paymentHandler.UpdatePayment)
}
