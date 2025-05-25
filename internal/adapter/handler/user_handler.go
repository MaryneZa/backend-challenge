package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MaryneZa/backend-challenge/internal/core/port"
	"github.com/MaryneZa/backend-challenge/internal/core/util"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.SendErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendErrorResponse(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if err := uh.userService.Register(r.Context(), req.Email, req.Password); err != nil {
		util.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendSuccessResponse(w, map[string]interface{}{
		"message": "Register successfully !",
	})
}

func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.SendErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ID string `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendErrorResponse(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	user, err := uh.userService.FindByID(r.Context(), req.ID)
	if err != nil {
		util.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendSuccessResponse(w, map[string]interface{}{
		"user": user,
	})

}

func (uh *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.SendErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendErrorResponse(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	user, err := uh.userService.FindByEmail(r.Context(), req.Email)
	if err != nil {
		util.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendSuccessResponse(w, map[string]interface{}{
		"user": user,
	})

}
func (uh *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.SendErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := uh.userService.GetAllUser(r.Context())
	if err != nil {
		util.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendSuccessResponse(w, map[string]interface{}{
		"user": users,
	})

}

// UpdateEmail(ctx context.Context, id bson.ObjectID, email string) error
func (uh *UserHandler) UpdateEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		util.SendErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendErrorResponse(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	userID, ok := util.GetUserID(w, r)
	if !ok {
		return
	}

	if err := uh.userService.UpdateEmail(r.Context(), userID, req.Email); err != nil {
		util.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendSuccessResponse(w, map[string]interface{}{
		"message": "Update successfully !",
	})

}

// UpdateName(ctx context.Context, id bson.ObjectID, name string) error
func (uh *UserHandler) UpdateName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		util.SendErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendErrorResponse(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	userID, ok := util.GetUserID(w, r)
	if !ok {
		return
	}

	if err := uh.userService.UpdateName(r.Context(), userID, req.Name); err != nil {
		util.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendSuccessResponse(w, map[string]interface{}{
		"message": "Update successfully !",
	})
}

func (uh *UserHandler) DeleteByEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		util.SendErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendErrorResponse(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if err := uh.userService.Delete(r.Context(), req.Email); err != nil {
		util.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendSuccessResponse(w, map[string]interface{}{
		"message": "Delete " + req.Email + " successfully !",
	})
}
