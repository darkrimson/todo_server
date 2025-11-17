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
		utils.WriteJSON(w, common.Response{Error: "invalid JSON"}, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if task.Title == "" {
		utils.WriteJSON(w, common.Response{Error: "title is required"}, http.StatusBadRequest)
		return
	}

	if err := utils.CheckDate(&task); err != nil {
		utils.WriteJSON(w, common.Response{Error: "invalid date"}, http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		utils.WriteJSON(w, common.Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, common.Response{ID: id}, http.StatusOK)
}
