package storage

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/L3HSec/trash-file-system/common"
	"github.com/L3HSec/trash-file-system/dbman"
	"github.com/L3HSec/trash-file-system/fileman"
)

type diskStorageManager struct {
	dbMan   common.DatabaseManager
	fileMan common.FileManager
}

func generateFileID() common.FileID {
	id := rand.Int63()
	return common.FileID(id)
}

func (p *diskStorageManager) SaveFile(fileName string, comment string, reader io.Reader) (*common.File, error) {
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

	fmt.Printf("Saving file %s id %X", file.FileName, file.ID)

	return &file, nil
}

func (p *diskStorageManager) GetFile(id common.FileID) (*common.File, io.ReadCloser, error) {
	file, err := p.dbMan.QueryFile(id)
	if err != nil {
		return nil, nil, common.NewError("no such file id in db").Base(err)
	}

	storedName := fmt.Sprintf("%X", file.ID)
	reader, err := p.fileMan.GetFile(storedName)
	if err != nil {
		return nil, nil, common.NewError("failed to open file").Base(err)
	}
	fmt.Printf("Getting file %s id %X", file.FileName, file.ID)
	return file, reader, nil
}

func (p *diskStorageManager) ListFiles() ([]common.File, error) {
	return p.dbMan.ListFiles()
}

func (p *diskStorageManager) AutoClean() {
	for {
		files, err := p.dbMan.ListFiles()
		if err != nil {
			panic("cannot list all files")
		}
		for _, f := range files {
			if time.Now().After(f.Expire) {
				p.dbMan.DeleteFile(f.ID)
				p.fileMan.DeleteFile(fmt.Sprintf("%X", f.ID))
				fmt.Println("Expired file deleted: " + f.FileName)
			}
		}
		time.Sleep(time.Minute * 1)
	}
}

//NewStorageManager creates storage manager
func NewStorageManager(dbPath string, uploadDir string) common.StorageManager {
	man := &diskStorageManager{
		dbMan:   dbman.NewManager(dbPath),
		fileMan: fileman.NewManager(uploadDir),
	}
	go man.AutoClean()
	return man
}
