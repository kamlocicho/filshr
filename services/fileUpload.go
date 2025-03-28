package services

import (
	"io"
	"net/http"
	"os"
)

// This will be integrated with external storage bucket in future
func FileUpload(r *http.Request, filePath string) (string, error) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	f, err := os.OpenFile(filePath+handler.Filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()

	io.Copy(f, file)
	return handler.Filename, nil
}
