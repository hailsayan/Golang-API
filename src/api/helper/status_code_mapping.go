package helper

import (
	"net/http"

	"github.com/hailsayan/Golang-API/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{

	// OTP
	service_errors.OptExists:   409,
	service_errors.OtpUsed:     409,
	service_errors.OtpNotValid: 400,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
