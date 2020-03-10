package common

import (
	"io"
	"time"
)

type FileID uint64

type File struct {
	FileName string
	FileSize int
	Comment  string
	ID       FileID
	Expire   time.Time
}

type StorageManager interface {
	SaveFile(fileName string, comment string, reader io.Reader) (*File, error)
	GetFile(id FileID) (io.ReadCloser, error)
	ListFiles() ([]File, error)
}

type DatabaseManager interface {
	AddFile(file *File) error
	QueryFile(id FileID) (*File, error)
	DeleteFile(id FileID) error
}
