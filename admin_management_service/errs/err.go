package errs

import "errors"

var (
	UnexpectedError       = errors.New("unexpected error")
	VideoDocNotFoundError = errors.New("video document not found")
	DocImageNotFoundError = errors.New("document's image not found")
	DrugDocNotFoundError  = errors.New("drug document not found")
)
