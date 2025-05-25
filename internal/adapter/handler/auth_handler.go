package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MaryneZa/backend-challenge/internal/core/port"
	"github.com/MaryneZa/backend-challenge/internal/core/util"

)

type AuthHandler struct {
	authService port.AuthService
}

func NewAuthHandler(authService port.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
} 

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.SendErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Email string `json:"email"`
		Password string `json:"password"`
 	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendErrorResponse(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	token, err := ah.authService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		util.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendSuccessResponse(w, map[string]interface{}{
		"token":   token,
	})
}