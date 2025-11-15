package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/username/go-final-project/pkg/db"
	"github.com/username/go-final-project/pkg/db/model"
	"github.com/username/go-final-project/pkg/utils"
)

type response struct {
	ID    int64  `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{Error: "invalid JSON"})
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{Error: "title is required"})
		return
	}

	if err := utils.CheckDate(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{Error: "invalid date"})
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{Error: err.Error()})
		return
	}

	utils.WriteJSON(w, response{ID: id}, http.StatusOK)
}
