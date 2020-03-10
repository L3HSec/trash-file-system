package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/L3HSec/trash-file-system/common"
	"github.com/L3HSec/trash-file-system/server"
	"github.com/gorilla/mux"
)

type apiManager struct {
	storageMan common.StorageManager
}

func (p *apiManager) handleUpload(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20)
	reader, handler, err := req.FormFile("upload")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer reader.Close()
	fileName := handler.Filename
	comment := req.FormValue("comment")

	file, err := p.storageMan.SaveFile(fileName, comment, reader)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	respBody, _ := json.Marshal(fileUploadResponse{
		FileID: fmt.Sprintf("%X", file.ID),
		Expire: file.Expire.Unix(),
	})
	w.WriteHeader(201)
	w.Write(respBody)
}

func (p *apiManager) handleDownload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileIDStr, found := vars["id"]
	if !found {
		w.WriteHeader(400)
		return
	}
	fmt.Println(fileIDStr)
	var fileID common.FileID
	_, err := fmt.Sscanf(fileIDStr, "%X", &fileID)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	file, fileReader, err := p.storageMan.GetFile(fileID)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer fileReader.Close()
	w.Header().Set("Content-Disposition", "attachment; filename="+file.FileName)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", file.FileSize))
	w.WriteHeader(200)
	io.Copy(w, fileReader)
}

func (p *apiManager) handleList(w http.ResponseWriter, r *http.Request) {
	files, err := p.storageMan.ListFiles()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	filesList := fileListResponse{
		files: make([]fileInfo, 0, len(files)),
	}

	for _, f := range files {
		filesList.files = append(filesList.files, fileInfo{
			FileName: f.FileName,
			FileSize: f.FileSize,
			Expire:   f.Expire.Unix(),
			FileID:   fmt.Sprintf("%X", f.ID),
			Comment:  f.Comment,
		})
	}

	respBody, err := json.Marshal(filesList.files)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(respBody)
}

//Run initializes apis
func Run(storageMan common.StorageManager) {
	man := apiManager{
		storageMan: storageMan,
	}
	server.RegisterAPI("POST", "/file", man.handleUpload)
	server.RegisterAPI("GET", "/file/{id:[0-9A-Fa-f]+}", man.handleDownload)
	server.RegisterAPI("GET", "/file", man.handleList)
}
