package controllers

import (
	"awesomeBackend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	carDetailsService services.CarDetailsService
)

type CarDetailsController interface {
	GetCarDetails(ctx *gin.Context)
}

func NewCarDetailsController(service services.CarDetailsService) CarDetailsController {
	carDetailsService = service
	return &controller{}
}

func (c *controller) GetCarDetails(ctx *gin.Context) {
	carDetails := carDetailsService.GetDetails()
	ctx.JSON(http.StatusOK, carDetails)
}
