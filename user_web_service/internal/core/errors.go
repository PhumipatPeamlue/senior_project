package core

type ErrNotFound struct {
	OriginalError error
}

func (e *ErrNotFound) Error() string {
	return e.OriginalError.Error()
}

type ErrDuplicate struct {
	OriginalError error
}

func (e *ErrDuplicate) Error() string {
	return e.OriginalError.Error()
}
