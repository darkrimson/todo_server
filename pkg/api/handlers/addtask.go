package handlers

import (
	"encoding/json"
	"net/http"

	"go1f/pkg/api/common"
	"go1f/pkg/db"
	"go1f/pkg/db/model"
	"go1f/pkg/utils"
)

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Error: "invalid JSON"})
		return
	}
	defer r.Body.Close()

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Error: "title is required"})
		return
	}

	if err := utils.CheckDate(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Error: "invalid date"})
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Error: err.Error()})
		return
	}

	utils.WriteJSON(w, common.Response{ID: id}, http.StatusOK)
}
