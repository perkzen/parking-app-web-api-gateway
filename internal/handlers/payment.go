package handlers

import (
	"github.com/gin-gonic/gin"
	payments "parking-app-web-api-gateway/internal/payments/proto"
	"parking-app-web-api-gateway/internal/services"
)

type PaymentHandler struct {
	PaymentService *services.PaymentService
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{
		PaymentService: services.NewPaymentService(),
	}
}

// UpdatePayment godoc
// @Tags payments
// @Summary Update a payment
// @Description Update a payment
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Param payment body payments.UpdatePaymentRequest true "Payment"
// @Success 200 {string} string
// @Router /payments/{id} [put]
func (p *PaymentHandler) UpdatePayment(c *gin.Context) {

	var updatePaymentRequest payments.UpdatePaymentRequest
	err := c.ShouldBindJSON(&updatePaymentRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	payment, err := p.PaymentService.UpdatePayment(&updatePaymentRequest)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, payment)
}

// GetPayments godoc
// @Tags payments
// @Summary Get all payments
// @Description Get all payments
// @Produce json
// @Success 200 {array} payments.GetAllPaymentsResponse
// @Router /payments [get]
func (p *PaymentHandler) GetPayments(c *gin.Context) {
	list, err := p.PaymentService.GetPayments()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, list)
}
