package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	key  = os.Getenv("USAGE_KEY")
	port = flag.String("p", "4000", "Port that the app listens to")
	dir  = flag.String("dir", "./uploads/", "Directory to store images in")
)

func init() {
	if key == "" {
		key = randString(10)
	}
	log.Printf("No key was specified... using [%s]. This is not recommended.", key)
}

func main() {
	r := mux.NewRouter()

	// Handle new uploads
	r.HandleFunc("/upload", handleUpload).
		Methods(http.MethodPost).
		Headers("Authorization", key)

	// Serve pre-existing images
	r.HandleFunc("/{image}", http.FileServer(http.Dir(*dir)).ServeHTTP).
		Methods(http.MethodGet)

	http.ListenAndServe(":"+*port, r)
}
