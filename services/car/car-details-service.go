package services

import (
	"awesomeBackend/entities"
	"github.com/goccy/go-json"
	"net/http"
)

type CarDetailsService interface {
	GetDetails() entities.CarDetails
}

var (
	carService       CarService   = NewCarService()
	ownerService     OwnerService = NewOwnerService()
	carDataChannel                = make(chan *http.Response)
	ownerDataChannel              = make(chan *http.Response)
)

type detailService struct{}

func NewCarDetailsService() CarDetailsService {
	return &detailService{}
}

func (s *detailService) GetDetails() entities.CarDetails {
	// goroutine call endpoint 1
	go carService.FetchData()

	// goroutine call endpoint 2
	go ownerService.FetchData()

	var carData, ownerData *http.Response

	for i := 0; i < 2; i++ {
		select {
		case carResponse := <-carDataChannel:
			carData = carResponse
		case ownerResponse := <-ownerDataChannel:
			ownerData = ownerResponse
		}
	}

	var car entities.Car
	var owner entities.Owner

	json.NewDecoder(carData.Body).Decode(&car)
	json.NewDecoder(ownerData.Body).Decode(&owner)

	return entities.CarDetails{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}

}
