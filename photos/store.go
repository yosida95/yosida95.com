package photos

import (
	"context"
	"errors"
	"io"
)

var ErrPhotoNotFound = errors.New("photo not found")

type StoreFactory interface {
	Close() error
	Begin(context.Context) (Store, error)
}

type Store interface {
	Rollback() error
	Commit() error

	FetchPhoto() PhotoFetcher
	CountPhoto() (int64, error)

	BlobGet(string) (io.ReadCloser, error)
}

type PhotoFetcher interface {
	Id(PhotoId) PhotoFetcher
	Limit(int64) PhotoFetcher
	Offset(int64) PhotoFetcher
	First() (*Photo, error)
	All() ([]*Photo, error)
}
