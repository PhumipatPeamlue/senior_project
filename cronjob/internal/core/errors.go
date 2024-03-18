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

type ErrPetNotFound struct {
	baseError
}

func NewErrPetNotFound(err error) error {
	return &ErrPetNotFound{
		baseError: baseError{originalError: err},
	}
}

type ErrReminderNotFound struct {
	baseError
}

func NewErrReminderNotFound(err error) error {
	return &ErrReminderNotFound{
		baseError: baseError{originalError: err},
	}
}

type ErrNotificationNotFound struct {
	baseError
}

func NewErrNotificationNotFound(err error) error {
	return &ErrNotificationNotFound{
		baseError: baseError{originalError: err},
	}
}
