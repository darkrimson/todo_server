// pkg/handlers/signin.go

package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"go1f/pkg/api/auth"
	"go1f/pkg/api/common"
	"go1f/pkg/utils"
)

type SigninRequest struct {
	Password string `json:"password"`
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	var req SigninRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSON(w, common.Response{Error: "invalid JSON"}, http.StatusUnauthorized)
		return
	}

	expectedPassword := os.Getenv("TODO_PASSWORD")
	if expectedPassword == "" {
		utils.WriteJSON(w, common.Response{Error: "empty password"}, http.StatusUnauthorized)
		return
	}

	if req.Password != expectedPassword {
		utils.WriteJSON(w, common.Response{Error: "wrong password"}, http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(req.Password)
	if err != nil {
		utils.WriteJSON(w, common.Response{Error: "internet server error"}, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
