package api

type fileUploadResponse struct {
	FileID string
	Expire int64
	Info   string
}

type fileUploadRequest struct {
	FileName string
	Comment  string
}

type fileInfo struct {
	FileName string
	FileSize int64
	Comment  string
	FileID   string
	Expire   int64
}

type fileListResponse struct {
	files []fileInfo
}
