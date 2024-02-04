package services

var (
	NotFoundError  error = notFoundError{}
	DuplicateError error = duplicateError{}
)

type notFoundError struct{}

func (e notFoundError) Error() string {
	return "not found error"
}

type duplicateError struct{}

func (e duplicateError) Error() string {
	return "duplicate error"
}
