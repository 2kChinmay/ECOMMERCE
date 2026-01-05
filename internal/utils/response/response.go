package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string
	Error  string
}

const (
	StatusOk    = "OK"
	StatusERROR = "ERROR"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := json.NewEncoder(w).Encode(data)
	return response
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusERROR,
		Error:  err.Error(),
	}
}

func ValidaionError(errs validator.ValidationErrors) Response {
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is a required field", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}

	return Response{
		Status: StatusERROR,
		Error:  strings.Join(errMsgs, ", "),
	}
}
