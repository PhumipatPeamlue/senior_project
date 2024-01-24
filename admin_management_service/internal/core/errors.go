package core

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"net/http"
)

const (
	CodeErrorNotFound = iota
	CodeErrorDuplicate
	CodeErrorBadRequest
)

type Error struct {
	originalError error
	code          int
}

func WrapError(err error, code int) error {
	return &Error{
		originalError: err,
		code:          code,
	}
}

func NewErrorNotFound(originalError error) error {
	return WrapError(originalError, CodeErrorNotFound)
}

func NewErrorDuplicate(originalError error) error {
	return WrapError(originalError, CodeErrorDuplicate)
}

func NewErrorBadRequest(originalError error) error {
	return WrapError(originalError, CodeErrorBadRequest)
}

func (e *Error) Error() string {
	return e.originalError.Error()
}

func (e *Error) Code() int {
	return e.code
}

func HandleEsErrorResponse(res *esapi.Response) error {
	code := res.StatusCode
	err := fmt.Errorf("%s", res.String())
	switch code {
	case http.StatusNotFound:
		return NewErrorNotFound(err)
	case http.StatusBadRequest:
		return NewErrorBadRequest(err)
	}
	return err
}
