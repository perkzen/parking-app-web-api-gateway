package handlers

import (
	"github.com/gin-gonic/gin"
	"parking-app-web-api-gateway/internal/services"
)

type ParkingHandler struct {
	ParkingService *services.ParkingService
}

func NewParkingHandler() *ParkingHandler {
	return &ParkingHandler{
		ParkingService: services.NewParkingService(),
	}
}

// GetParkingSpots godoc
// @Tags parking
// @Summary Get parking spots
// @Description Get parking spots
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} services.ParkingSpot
// @Router /park [get]
func (p *ParkingHandler) GetParkingSpots(c *gin.Context) {
	parkingSpots, err := p.ParkingService.GetParkingSpots()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, parkingSpots)
}

// GetParkingSpot godoc
// @Tags parking
// @Summary Get parking spot
// @Description Get parking spot
// @Produce json
// @Security ApiKeyAuth
// @Param name path string true "Parking spot name"
// @Success 200 {object} services.ParkingSpot
// @Router /park/{name} [get]
func (p *ParkingHandler) GetParkingSpot(c *gin.Context) {
	name := c.Param("name")

	parkingSpot, err := p.ParkingService.GetParkingSpot(name)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Not found",
		})
		return
	}

	c.JSON(200, parkingSpot)
}

// CreateParkingSpot godoc
// @Tags parking
// @Summary Create parking spot
// @Description Create parking spot
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body services.CreateParkingSpotRequest true "Create parking spot request"
// @Success 200 {object} services.ParkingSpot
// @Router /park [post]
func (p *ParkingHandler) CreateParkingSpot(c *gin.Context) {
	var request services.CreateParkingSpotRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := p.ParkingService.AddParkingSpot(&request)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Parking spot created",
	})
}
