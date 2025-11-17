package handlers

import (
	"net/http"
	"time"

	"go1f/pkg/api/common"
	"go1f/pkg/db"
	"go1f/pkg/utils"
)

func DoneTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		utils.WriteJSON(w, common.Response{Error: "id is required"}, http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		utils.WriteJSON(w, common.Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			utils.WriteJSON(w, common.Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}
		utils.WriteJSON(w, common.Response{}, http.StatusOK)
		return
	}

	now := time.Now()
	nextDate, err := utils.NextDate(now, task.Date, task.Repeat)
	if err != nil {
		utils.WriteJSON(w, common.Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	task.Date = nextDate
	err = db.UpdateTask(task)
	if err != nil {
		utils.WriteJSON(w, common.Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, common.Response{}, http.StatusOK)
}
