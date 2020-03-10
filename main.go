package main

import (
	"github.com/L3HSec/trash-file-system/api"
	"github.com/L3HSec/trash-file-system/common"
	"github.com/L3HSec/trash-file-system/server"
	"github.com/L3HSec/trash-file-system/storage"
)

func main() {
	var storageMan common.StorageManager
	storageMan = storage.NewStorageManager("tfs.db", "upload_files/")
	api.Run(storageMan)
	server.Run(":8080")
}
