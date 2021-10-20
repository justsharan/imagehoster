package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	key  = os.Getenv("USAGE_KEY")
	port = flag.String("p", "4000", "Port that the app listens to")
	dir  = flag.String("dir", "./uploads/", "Directory to store images in")
)

func main() {
	r := mux.NewRouter()

	// Serve pre-existing images
	r.HandleFunc("/{image}", http.FileServer(http.Dir(*dir)).ServeHTTP).
		Methods(http.MethodGet)

	http.ListenAndServe(":"+*port, r)
}
