package server

import (
	"net/http"
	"os"

	"github.com/username/go-final-project/pkg/api"
)

func Run() error {
	port := os.Getenv("TODO_PORT")

	if port == "" {
		port = "7540"
	}
	webDir := http.FileServer(http.Dir("web"))

	http.Handle("/", webDir)
	api.Init()
	return http.ListenAndServe(":"+port, nil)
}
