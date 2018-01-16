package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createRouter(staticPath string) *mux.Router {
	log.Print("Creating router using static path: ", staticPath)
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir(staticPath))
	router.PathPrefix("/s/").Handler(http.StripPrefix("/s/", fs))

	router.Handle("/uptime", newUptimeHandler())

	return router
}
