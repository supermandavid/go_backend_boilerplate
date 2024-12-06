package controllers

import (
	"github.com/gin-gonic/gin"
	c "github.com/supermandavid/go_backend_boilerplate/services/car"
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
