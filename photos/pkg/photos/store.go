package photos

import (
	"errors"
	"io"
)

var ErrPhotoNotFound = errors.New("photo not found")

type Store interface {
	Rollback() error
	Commit() error

	PhotoFetch() PhotoFetcher
	PhotoCount() (int64, error)
	// PhotoNextId() (PhotoId, error)
	// PhotoSave(*Photo) error

	BlobGet(string) (io.ReadCloser, error)
}

type PhotoFetcher interface {
	Id(PhotoId) PhotoFetcher
	Limit(int64) PhotoFetcher
	Offset(int64) PhotoFetcher
	First() (*Photo, error)
	All() ([]*Photo, error)
}
