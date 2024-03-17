package core

type baseError struct {
	originalError error
}

func (e *baseError) Error() string {
	return e.originalError.Error()
}

type ErrUserNotFound struct {
	baseError
}

func NewErrUserNotFound(err error) error {
	return &ErrUserNotFound{
		baseError: baseError{originalError: err},
	}
}

type ErrUserDuplicate struct {
	baseError
}

func NewErrUserDuplicate(err error) error {
	return &ErrUserDuplicate{
		baseError: baseError{originalError: err},
	}
}
