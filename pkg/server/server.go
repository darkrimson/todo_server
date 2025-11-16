package server

import (
	"net/http"
	"os"

	"go1f/pkg/api"
)

func Run() error {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	r := api.InitRoute()

	// подключаем web-директорию к chi, а не к http.DefaultServeMux
	r.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("web"))))

	return http.ListenAndServe(":"+port, r)
}
