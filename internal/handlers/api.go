package handlers

import (
	"encoding/json"
	"fmt"
	"sword-project/internal/services"

	"github.com/go-playground/validator/v10"
)

type ApiHandler struct {
	Services services.Services
}

func NewApiHandler(services services.Services) *ApiHandler {
	return &ApiHandler{Services: services}
}

func buildValidationErrorMessage(err error) string {
	msg := err.Error()

	switch errType := err.(type) {

	case *json.UnmarshalTypeError:
		msg = errType.Field + " type error."

	case validator.ValidationErrors:
		msg = fmt.Sprintf("%s is %s", errType[0].Field(), errType[0].ActualTag())

	}

	return msg
}
