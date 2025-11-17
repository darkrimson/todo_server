package handlers

import (
	"encoding/json"
	"net/http"

	"go1f/pkg/api/common"
	"go1f/pkg/db"
	"go1f/pkg/db/model"
	"go1f/pkg/utils"
)

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(common.Response{Error: "invalid JSON"})
		return
	}
	defer r.Body.Close()

	id := task.ID
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(common.Response{Error: "id is required"})
		return
	}

	if task.ID != id {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(common.Response{Error: "id in URL does not match id in request body"})
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(common.Response{Error: "title is required"})
		return
	}

	if err := utils.CheckDate(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(common.Response{Error: "invalid date"})
		return
	}

	_, err := db.GetTask(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(common.Response{Error: "task not found"})
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(common.Response{Error: "internal server error"})
		return
	}

	utils.WriteJSON(w, nil, http.StatusOK)
}
