package store

import (
	"context"

	"github.com/yosida95/yosida95.com/photos/pkg/photos"
)

type StoreFactory interface {
	Close() error
	Begin(context.Context) (Store, error)
}

type Store interface {
	photos.Store
}
