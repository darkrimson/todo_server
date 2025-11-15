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

	r := api.InitRoute()

	// подключаем web-директорию к chi, а не к http.DefaultServeMux
	r.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("web"))))

	return http.ListenAndServe(":"+port, r)
}
