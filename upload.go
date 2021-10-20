package main

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type")[0:5] != "image" {
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}

	// Resolve the appropriate extension type
	ext, err := mime.ExtensionsByType(r.Header.Get("Content-Type"))
	if err != nil {
		http.Error(w, "Error resolving file extension", http.StatusInternalServerError)
		return
	}

	// Assign a new filename to this image
	filename := randString(5) + ext[len(ext)-1]

	// Create the file in the directory
	file, err := os.Create(path.Join(*dir, filename))
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}

	// Copy image to file
	io.Copy(file, r.Body)
	fmt.Fprintf(w, filename)
}
