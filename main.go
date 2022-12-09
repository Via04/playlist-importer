package main
// Test package to work with Google API

import (
	"log"
	"net/http"

	"github.com/via04/playlist_importer/internal/handlers"
)

func main() {
	// Server Entry Point
	server := &http.Server{
		Addr: "127.0.0.1:1667",
		Handler: handlers.New(),
	}
	log.Fatal(server.ListenAndServe())
}