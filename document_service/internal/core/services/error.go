package services

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"net/http"
)

var (
	NotFoundError   error = notFoundError{}
	DuplicateError  error = duplicateError{}
	BadRequestError error = badRequestError{}
)

type notFoundError struct{}

func (e notFoundError) Error() string {
	return "not found error"
}

type duplicateError struct{}

func (e duplicateError) Error() string {
	return "duplicate error"
}

type badRequestError struct{}

func (e badRequestError) Error() string {
	return "bad request error"
}

func HandleEsErrorResponse(res *esapi.Response) error {
	code := res.StatusCode
	err := fmt.Errorf("%s", res.String())
	switch code {
	case http.StatusNotFound:
		return NotFoundError
	case http.StatusBadRequest:
		return BadRequestError
	}
	return err
}
