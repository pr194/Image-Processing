package controllers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	_, err = file.(io.Seeker).Seek(0, io.SeekStart)
	if err != nil {
		http.Error(w, "Failed to reset file reader", http.StatusInternalServerError)
		return
	}
	contentType := http.DetectContentType(buffer)
	allowedTypes := []string{"image/jpeg", "image/png"}
	isValidType := false
	for _, allowedType := range allowedTypes {
		if strings.HasPrefix(contentType, allowedType) {
			isValidType = true
			break
		}
	}
	if !isValidType {
		http.Error(w, "Invalid file type. Only JPEG and PNG are allowed.", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "File '%s' uploaded successfully with type: %s\n", header.Filename, contentType)

}
