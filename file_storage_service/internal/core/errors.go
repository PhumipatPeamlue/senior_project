package core

type baseError struct {
	originalError error
}

func (b *baseError) Error() string {
	return b.originalError.Error()
}

type ErrFileInfoNotFound struct {
	baseError
}

func NewErrFileInfoNotFound(err error) error {
	return &ErrFileInfoNotFound{
		baseError{originalError: err},
	}
}

type ErrFileInfoDuplicate struct {
	baseError
}

func NewErrFileInfoDuplicate(err error) error {
	return &ErrFileInfoDuplicate{
		baseError{originalError: err},
	}
}
