package services

import (
	"fmt"
	"net/http"
)

type CarService interface {
	FetchData()
}

const (
	carServiceUrl = "https://myfakeapi.com/api/cars/1"
)

type FetchCarDataService struct {
}

func NewCarService() CarService {
	return &FetchCarDataService{}
}

func (service *FetchCarDataService) FetchData() {
	client := &http.Client{}
	fmt.Printf("\nFetching Car data from %s", carServiceUrl)

	// call the external API
	resp, _ := client.Get(carServiceUrl)

	// Write response to the channel
	carDataChannel <- resp
}
