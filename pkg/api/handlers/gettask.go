package handlers

import (
	"net/http"

	"go1f/pkg/api/common"
	"go1f/pkg/db/repo"
	"go1f/pkg/utils"
)

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		utils.WriteJSON(w, common.Response{Error: "id is required"}, http.StatusBadRequest)
		return
	}

	task, err := repo.GetTask(id)
	if err != nil {
		utils.WriteJSON(w, common.Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, task, http.StatusOK)
}
