package handlers

import (
	"net/http"

	"github.com/username/go-final-project/pkg/db"
	"github.com/username/go-final-project/pkg/utils"
)

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, `{"error":"id is required"}`, http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, task, http.StatusOK)
}
