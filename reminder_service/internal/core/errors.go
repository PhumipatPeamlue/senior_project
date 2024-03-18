package core

type baseError struct {
	originalError error
}

func (e *baseError) Error() string {
	return e.originalError.Error()
}

// reminder service error

type ErrReminderNotFound struct {
	baseError
}

func NewErrReminderNotFound(err error) error {
	return &ErrReminderNotFound{
		baseError: baseError{originalError: err},
	}
}

type ErrReminderDuplicate struct {
	baseError
}

func NewErrReminderDuplicate(err error) error {
	return &ErrReminderDuplicate{
		baseError: baseError{originalError: err},
	}
}

// notification service error

type ErrNotificationStatusConflict struct {
	baseError
}

func NewErrNotificationStatusConflict(err error) error {
	return &ErrNotificationStatusConflict{
		baseError: baseError{originalError: err},
	}
}

type ErrNotificationStatusInternalServerError struct {
	baseError
}

func NewErrNotficationStatusInternalServerError(err error) error {
	return &ErrNotificationStatusInternalServerError{
		baseError: baseError{originalError: err},
	}
}
