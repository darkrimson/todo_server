package handlers

import (
	"net/http"
	"strconv"

	"go1f/pkg/api/common"
	"go1f/pkg/db/model"
	"go1f/pkg/db/repo"
	"go1f/pkg/utils"
)

type TasksResp struct {
	Tasks []*model.Task `json:"tasks"`
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")

	limit := 50
	if lim := r.URL.Query().Get("limit"); lim != "" {
		if parsed, err := strconv.Atoi(lim); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	tasks, err := repo.GetTasks(limit, search)
	if err != nil {
		utils.WriteJSON(w, common.Response{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	resp := TasksResp{Tasks: tasks}
	utils.WriteJSON(w, resp, http.StatusOK)
}
