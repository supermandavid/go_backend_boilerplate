package controllers

import (
	c "awesomeBackend/services/car"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	carDetailsService c.CarDetailsService
)

type CarDetailsController interface {
	GetCarDetails(ctx *gin.Context)
}

func NewCarDetailsController(service c.CarDetailsService) CarDetailsController {
	carDetailsService = service
	return &controller{}
}

func (c *controller) GetCarDetails(ctx *gin.Context) {
	carDetails := carDetailsService.GetDetails()
	ctx.JSON(http.StatusOK, carDetails)
}
