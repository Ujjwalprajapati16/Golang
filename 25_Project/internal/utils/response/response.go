package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json: "status"`
	Error  string `json: "error"`
}

const (
	StatusOk    = "OK"
	StatusError = "ERROR"
)

// WriteJson writes a json response
func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// GeneralErro returns a general error
func GeneralErro(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

// validate error
func ValidateError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, err.Field()+" is required")
		default:
			errMsgs = append(errMsgs, "invalid "+err.Field())
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
