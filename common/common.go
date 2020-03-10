package common

import (
	"io"
	"time"
)

type FileID uint64

type File struct {
	FileName string
	FileSize int64
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
	ListFiles() ([]File, error)
	DeleteFile(id FileID) error
}

type FileManager interface {
	SaveFile(fileName string, reader io.Reader) (int64, error)
	GetFile(fileName string) (io.ReadCloser, error)
	DeleteFile(fileName string) error
}
