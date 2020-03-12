package fileman

import (
	"io"
	"os"

	"github.com/L3HSec/trash-file-system/common"
)

type fileManager struct {
	path string
}

func (p *fileManager) SaveFile(fileName string, reader io.Reader) (int64, error) {
	file, err := os.OpenFile(p.path+"/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}
	read, err := io.Copy(file, reader)
	return read, err
}

func (p *fileManager) GetFile(fileName string) (io.ReadCloser, error) {
	file, err := os.OpenFile(p.path+fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	return file, err
}

func (p *fileManager) DeleteFile(fileName string) error {
	return os.Remove(p.path + fileName)
}

//NewManager creates new file manager
func NewManager(path string) common.FileManager {
	os.Mkdir(path, 0766)
	return &fileManager{
		path: path,
	}
}
