package dbman

import (
	"database/sql"
	"time"

	"github.com/L3HSec/trash-file-system/common"
	_ "github.com/mattn/go-sqlite3"
)

type databaseManager struct {
	database *sql.DB
	common.DatabaseManager
}

func (p *databaseManager) AddFile(file *common.File) error {
	statement, _ := p.database.Prepare("INSERT INTO files (filename, filesize, id, expire, comment) VALUES (?,?,?,?,?)")
	_, err := statement.Exec(
		file.FileName,
		file.FileSize,
		file.ID,
		file.Expire.Unix(),
		file.Comment,
	)
	return err
}

func (p *databaseManager) QueryFile(id common.FileID) (*common.File, error) {
	rows, err := p.database.Query("SELECT * FROM files WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	var fileName, comment string
	var fileSize, expire int64
	var pk int
	if !rows.Next() {
		return nil, common.NewError("file id not found")
	}
	err = rows.Scan(&pk, &fileName, &fileSize, &id, &expire, &comment)
	if err != nil {
		return nil, err
	}
	return &common.File{
		FileName: fileName,
		FileSize: fileSize,
		ID:       id,
		Expire:   time.Unix(expire, 0),
		Comment:  comment,
	}, nil
}

func (p *databaseManager) DeleteFile(id common.FileID) error {
	_, err := p.database.Exec("DELETE FROM files WHERE id=?", id)
	return err
}

func (p *databaseManager) ListFiles() ([]common.File, error) {
	rows, err := p.database.Query("SELECT * FROM files")
	if err != nil {
		return nil, err
	}
	fileList := make([]common.File, 0, 128)
	for rows.Next() {
		var fileName, comment string
		var fileSize, id, expire int64
		var pk int
		err = rows.Scan(&pk, &fileName, &fileSize, &id, &expire, &comment)
		if err != nil {
			return nil, err
		}
		fileList = append(fileList, common.File{
			FileName: fileName,
			FileSize: fileSize,
			ID:       common.FileID(id),
			Expire:   time.Unix(expire, 0),
			Comment:  comment,
		})
	}
	return fileList, nil
}

func NewManager(path string) common.DatabaseManager {
	var err error
	database, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS files (pk INTEGER PRIMARY KEY, filename TEXT, filesize INT64, id INT64, expire INT64, comment TEXT)")
	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}

	return &databaseManager{
		database: database,
	}
}
