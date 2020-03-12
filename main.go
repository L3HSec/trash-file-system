package main

import (
	"fmt"

	"github.com/L3HSec/trash-file-system/api"
	"github.com/L3HSec/trash-file-system/common"
	"github.com/L3HSec/trash-file-system/conf"
	"github.com/L3HSec/trash-file-system/server"
	"github.com/L3HSec/trash-file-system/storage"
)

func main() {
	var storageMan common.StorageManager
	config, err := conf.GetConfig("config.json")
	if err != nil {
		fmt.Println("Failed to load config.json")
		return
	}
	fmt.Println("Config: \n", config)
	storageMan = storage.NewStorageManager(config["database"].(string), config["files_dir"].(string))
	api.Run(storageMan)
	server.Run(config["address"].(string))
}
