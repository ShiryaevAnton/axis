package apierrors

import (
	"fmt"
	"net/http"
)

//APIErr ...
type APIErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type apiErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e apiErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e apiErr) Message() string {
	return e.ErrMessage
}

func (e apiErr) Status() int {
	return e.ErrStatus
}

func (e apiErr) Causes() []interface{} {
	return e.ErrCauses
}

//APIError ...
func APIError(message string, status int, err string, causes []interface{}) APIErr {
	return apiErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

//BadRequestAPIError ..
func BadRequestAPIError(message string) APIErr {
	return apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

//NotFoundAPIError ...
func NotFoundAPIError(message string) APIErr {
	return apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

//UnauthorizedAPIError ...
func UnauthorizedAPIError(message string) APIErr {
	return apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

//InternalServerAPIError ...
func InternalServerAPIError(message string, err error) APIErr {
	result := apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}
