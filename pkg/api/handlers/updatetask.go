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
		utils.WriteJSON(w, common.Response{Error: "invalid JSON"}, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id := task.ID
	if id == "" {
		utils.WriteJSON(w, common.Response{Error: "id is required"}, http.StatusBadRequest)
		return
	}

	if task.ID != id {
		utils.WriteJSON(w, common.Response{Error: "id in URL does not match id in request body"}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		utils.WriteJSON(w, common.Response{Error: "title is required"}, http.StatusBadRequest)
		return
	}

	if err := utils.CheckDate(&task); err != nil {
		utils.WriteJSON(w, common.Response{Error: "invalid date"}, http.StatusBadRequest)
		return
	}

	_, err := db.GetTask(id)
	if err != nil {
		utils.WriteJSON(w, common.Response{Error: "task not found"}, http.StatusNotFound)
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		utils.WriteJSON(w, common.Response{Error: "internal server error"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, common.Response{}, http.StatusOK)
}
