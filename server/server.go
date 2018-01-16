package server

import (
	"fmt"
	"log"
	"net/http"
)

// ListenAndServe starts a HTTP server for the app
func ListenAndServe(port int, staticPath string) error {
	addr := fmt.Sprintf(":%d", port)
	router := createRouter(staticPath)

	log.Printf("Serving on %s...", addr)
	return http.ListenAndServe(addr, router)
}
