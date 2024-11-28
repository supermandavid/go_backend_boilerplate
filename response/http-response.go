package response

import (
	"fmt"
)

type HTTPResponse struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func NewHTTPResponse(status bool, message interface{}, data interface{}) HTTPResponse {
	// Set default values
	defaultMessage := "Operation completed"
	if message != nil {
		defaultMessage = fmt.Sprintf("%v", message)
	}


	return HTTPResponse{
		Message: defaultMessage,
		Status:  status,
		Data:    data,
	}
}
