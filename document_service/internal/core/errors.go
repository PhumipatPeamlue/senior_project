package core

type baseError struct {
	originalError error
}

func (b *baseError) Error() string {
	return b.originalError.Error()
}

type ErrDocNotFound struct {
	baseError
}

func NewErrDocNotFound(err error) error {
	return &ErrDocNotFound{baseError{originalError: err}}
}

type ErrDocDuplicate struct {
	baseError
}

func NewErrDocDuplicate(err error) error {
	return &ErrDocDuplicate{baseError{originalError: err}}
}

type ErrDocBadRequest struct {
	baseError
}

func NewErrDocBadRequest(err error) error {
	return &ErrDocBadRequest{baseError{originalError: err}}
}
