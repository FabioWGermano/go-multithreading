package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrInvalidJSON  = errors.New("invalid JSON")
)

type Error struct {
	statusCode int
	Errors     []string `json:"errors"`
}

func NewError(err error, status int) *Error {
	return &Error{
		statusCode: status,
		Errors:     []string{err.Error()},
	}
}

func NewErrorMessage(messages []string, status int) *Error {
	return &Error{
		statusCode: status,
		Errors:     messages,
	}
}

func (e *Error) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.statusCode)
	_ = json.NewEncoder(w).Encode(e)
}
