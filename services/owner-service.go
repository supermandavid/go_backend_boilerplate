package services

import (
	"fmt"
	"net/http"
)

type OwnerService interface {
	FetchData()
}

const (
	ownerServiceUrl = "https://myfakeapi.com/api/users/1"
)

type FetchOwnerDataService struct {
}

func NewOwnerService() OwnerService {
	return &FetchOwnerDataService{}
}

func (service *FetchOwnerDataService) FetchData() {
	client := &http.Client{}
	fmt.Printf("\nFetching Car data from %s", ownerServiceUrl)

	// call the external API
	resp, _ := client.Get(ownerServiceUrl)

	// Write response to the channel
	ownerDataChannel <- resp

}
