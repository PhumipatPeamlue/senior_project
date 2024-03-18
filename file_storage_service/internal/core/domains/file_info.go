package domains

import (
	"fmt"
	"net/url"
	"os"
)

type FileInfo struct {
	id       string
	fileName string
}

func (f *FileInfo) ID() string {
	return f.id
}

func (f *FileInfo) FileName() string {
	return f.fileName
}

func (f *FileInfo) URL(bucketName string) string {
	path := fmt.Sprintf("%s/%s/%s", "image", bucketName, f.fileName)
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	_url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", host, port),
		Path:   path,
	}

	return _url.String()
}

func (f *FileInfo) ChangeFileName(originalFileName string) {
	fileName := fmt.Sprintf("%s-%s", f.id, originalFileName)
	f.fileName = fileName
}

func ScanFileInfo(id, fileName string) FileInfo {
	return FileInfo{
		id:       id,
		fileName: fileName,
	}
}

func NewFileInfo(id, originalFileName string) FileInfo {
	fileName := fmt.Sprintf("%s-%s", id, originalFileName)
	return ScanFileInfo(id, fileName)
}
