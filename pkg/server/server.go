package server

import (
	"net/http"
	"os"
)

func Run() error {
	port := os.Getenv("TODO_PORT")

	if port == "" {
		port = "7540"
	}
	webDir := http.FileServer(http.Dir("web"))

	http.Handle("/", webDir)

	return http.ListenAndServe(":"+port, nil)
}
