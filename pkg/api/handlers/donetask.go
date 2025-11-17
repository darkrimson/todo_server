package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"go1f/pkg/api/common"
	"go1f/pkg/db"
	"go1f/pkg/utils"
)

func DoneTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(common.Response{Error: "id is required"})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(common.Response{Error: err.Error()})
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(common.Response{Error: err.Error()})
			return
		}
		utils.WriteJSON(w, common.Response{}, http.StatusOK)
		return
	}

	now := time.Now()
	nextDate, err := utils.NextDate(now, task.Date, task.Repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(common.Response{Error: err.Error()})
		return
	}

	task.Date = nextDate
	err = db.UpdateTask(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(common.Response{Error: err.Error()})
		return
	}

	utils.WriteJSON(w, common.Response{}, http.StatusOK)
}
