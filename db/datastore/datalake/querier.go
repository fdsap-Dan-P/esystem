package db

type QuerierLocal interface {
}

var _ QuerierLocal = (*QueriesLocal)(nil)
