package handlers

import (
	"net/http"

	"go1f/pkg/api/common"
	"go1f/pkg/db"
	"go1f/pkg/utils"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		utils.WriteJSON(w, common.Response{Error: "id is required"}, http.StatusBadRequest)
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		utils.WriteJSON(w, common.Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, common.Response{}, http.StatusOK)
}
