package handlers

import (
	"fmt"
	"github.com/echernukha/gocourse/simplevideoserver/tools"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const dirPath = "content"
const videoMimeType = "video/mp4"

func uploadVideo(w http.ResponseWriter, r *http.Request) {
	fileReader, header, err := r.FormFile("file[]")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	contentType := header.Header.Get("Content-Type")
	if contentType != videoMimeType {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id, err := uuid.NewRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := filepath.Join(id.String(), header.Filename)

	file, err := createFile(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tools.CloseQuietly(file)
	_, err = io.Copy(file, fileReader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprint(w, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createFile(fileName string) (*os.File, error) {
	filePath := filepath.Join(dirPath, fileName)
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return nil, err
	}
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
}
