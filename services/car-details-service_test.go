package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	carDetailsService = NewCarDetailsService()
)

func TestGetDetails(t *testing.T) {
	carDetails := carDetailsService.GetDetails()

	assert.NotNil(t, carDetails)
	assert.Equal(t, 1, carDetails.ID)
}
