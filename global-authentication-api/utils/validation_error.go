package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ValidationErrorResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Errors  map[string]interface{} `json:"errors"`
}

func CreateValidationErrorResponse(err error) ValidationErrorResponse {
	validationErrors := err.(validator.ValidationErrors)
	errorsMap := make(map[string]interface{})

	for _, fieldError := range validationErrors {
		errorsMap[fieldError.Field()] = fieldError.Tag()
	}

	return ValidationErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Validation failed",
		Errors:  errorsMap,
	}
}

func HandleError(w http.ResponseWriter, message string, statusCode int, errors map[string]interface{}) {
	response := ValidationErrorResponse{
		Code:    statusCode,
		Message: message,
		Errors:  errors,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
