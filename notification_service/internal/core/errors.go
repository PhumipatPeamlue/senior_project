package core

type baseError struct {
	originalError error
}

func (b *baseError) Error() string {
	return b.originalError.Error()
}

type ErrNotificationDuplicate struct {
	baseError
}

func NewErrNotificationDuplicate(err error) error {
	return &ErrNotificationDuplicate{
		baseError: baseError{originalError: err},
	}
}
