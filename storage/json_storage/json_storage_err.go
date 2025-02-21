package json_storage

import (
	"errors"
)

var (
	ErrStorageFileRead   = errors.New("failed to read storage file")
	ErrUnmarshalOrders   = errors.New("failed to unmarshal orders")
	ErrMarshalOrders     = errors.New("failed to marshal orders")
	ErrStorageFileWrite  = errors.New("failed to write storage file")
	ErrStorageFileCreate = errors.New("failed to create storage file")
	ErrStorageSave       = errors.New("failed to save storage")
	ErrStorageLoad       = errors.New("failed to load storage")
)
