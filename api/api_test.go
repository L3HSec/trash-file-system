package api

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/L3HSec/trash-file-system/common"
)

type mockStorageManager struct{}

func (p *mockStorageManager) SaveFile(fileName string, comment string, reader io.Reader) (*common.File, error) {
	fmt.Println("name: " + fileName)
	fmt.Println("comment: " + comment)
	buffer, _ := ioutil.ReadAll(reader)
	fmt.Println("content: " + string(buffer))
	return &common.File{
		ID:     123456,
		Expire: time.Now(),
	}, nil
}

func (p *mockStorageManager) GetFile(id common.FileID) (io.ReadCloser, error) {
	buffer := []byte{0x41, 0x42, 0x43, 0x44, 0x45, 0x46}
	return ioutil.NopCloser(bytes.NewReader(buffer)), nil
}

func (p *mockStorageManager) ListFiles() ([]common.File, error) {
	return []common.File{common.File{
		FileName: "cyka blayt",
		FileSize: 2333,
		ID:       0xdeadbeef,
		Comment:  "wdnmd",
		Expire:   time.Now(),
	}}, nil
}
