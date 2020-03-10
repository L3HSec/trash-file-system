package common

import (
	"io"
	"time"
)

//FileID unique id for files
type FileID uint64

//File struct describes a uploaded file
type File struct {
	FileName string
	FileSize int64
	Comment  string
	ID       FileID
	Expire   time.Time
}

//StorageManager manage the whole storage system
type StorageManager interface {
	SaveFile(fileName string, comment string, reader io.Reader) (*File, error)
	GetFile(id FileID) (*File, io.ReadCloser, error)
	ListFiles() ([]File, error)
}

//DatabaseManager handle the database
type DatabaseManager interface {
	AddFile(file *File) error
	QueryFile(id FileID) (*File, error)
	ListFiles() ([]File, error)
	DeleteFile(id FileID) error
}

//FileManager handle the uploaded files, notice that the filename is the real name of the file
type FileManager interface {
	SaveFile(fileName string, reader io.Reader) (int64, error)
	GetFile(fileName string) (io.ReadCloser, error)
	DeleteFile(fileName string) error
}
