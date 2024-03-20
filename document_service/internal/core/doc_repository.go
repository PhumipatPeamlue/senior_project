package core

type DocRepositoryInterface interface {
	Pagination(query string) (docs []any, total int, err error)
}
