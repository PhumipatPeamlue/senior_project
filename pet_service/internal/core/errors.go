package core

type baseError struct {
	originalError error
}

func (e *baseError) Error() string {
	return e.originalError.Error()
}

type ErrPetNotFound struct {
	baseError
}

func NewErrPetNotFound(err error) error {
	return &ErrPetNotFound{
		baseError: baseError{originalError: err},
	}
}

type ErrPetDuplicate struct {
	baseError
}

func NewErrPetDuplicate(err error) error {
	return &ErrPetDuplicate{
		baseError: baseError{originalError: err},
	}
}
