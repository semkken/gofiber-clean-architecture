package storage

import (
	"os"
)

// InitializeStorage creates the uploads directory if it doesn't exist
func InitializeStorage() error {
	return os.MkdirAll("./internal/storage/uploads", os.ModePerm)
}
