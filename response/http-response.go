package response

import (
	"encoding/json"
	"fmt"
)

type HTTPResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Data    []byte `json:"data"`
}

func NewHTTPResponse(message interface{}, status *bool, data interface{}) HTTPResponse {
	// Set default values
	defaultMessage := "Operation completed"
	if message != nil {
		defaultMessage = fmt.Sprintf("%v", message)
	}

	defaultStatus := true
	if status != nil {
		defaultStatus = *status
	}

	var dataBytes []byte
	switch v := data.(type) {
	case nil:
		dataBytes = []byte("null")
	case string:
		dataBytes = []byte(v)
	case []byte:
		dataBytes = v
	default:
		// Assume it's a list of structs or any other type; marshal to JSON
		jsonData, err := json.Marshal(v)
		if err != nil {
			dataBytes = []byte("null")
		} else {
			dataBytes = jsonData
		}
	}

	return HTTPResponse{
		Message: defaultMessage,
		Status:  defaultStatus,
		Data:    dataBytes,
	}
}
