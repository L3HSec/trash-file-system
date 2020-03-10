package storage

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/L3HSec/trash-file-system/common"
	"github.com/L3HSec/trash-file-system/dbman"
	fileman "github.com/L3HSec/trash-file-system/fileman"
)

type DiskStorageManager struct {
	dbMan   common.DatabaseManager
	fileMan common.FileManager
}

var Manager common.StorageManager

func generateFileID() common.FileID {
	id := rand.Int63()
	return common.FileID(id)
}

func (p *DiskStorageManager) SaveFile(fileName string, comment string, reader io.Reader) (*common.File, error) {
	file := common.File{
		FileName: fileName,
		ID:       generateFileID(),
		Comment:  comment,
		Expire:   time.Now().Add(time.Minute * 15),
	}

	storedName := fmt.Sprintf("%X", file.ID)

	fileSize, err := p.fileMan.SaveFile(storedName, reader)

	if err != nil {
		return nil, common.NewError("failed to save the file").Base(err)
	}

	file.FileSize = fileSize

	err = p.dbMan.AddFile(&file)

	if err != nil {
		return nil, common.NewError("failed to add file to db").Base(err)
	}

	return &file, nil
}

func (p *DiskStorageManager) GetFile(id common.FileID) (io.ReadCloser, error) {
	file, err := p.dbMan.QueryFile(id)

	if err != nil {
		return nil, common.NewError("no such file id in db").Base(err)
	}

	storedName := fmt.Sprintf("%X", file.ID)

	reader, err := p.fileMan.GetFile(storedName)

	if err != nil {
		return nil, common.NewError("failed to open file").Base(err)
	}

	return reader, nil
}

func (p *DiskStorageManager) ListFiles() ([]common.File, error) {
	return p.dbMan.ListFiles()
}

func init() {
	Manager = &DiskStorageManager{
		dbMan:   dbman.NewManager(dbPath),
		fileMan: fileman.NewManager(uploadDir),
	}
}
