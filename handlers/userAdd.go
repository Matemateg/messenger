package handlers

import (
	"encoding/json"
	"messenger/service"
	"net/http"
)

type AddingUser struct {
	service *service.UserService
}

func NewAddingUser(service *service.UserService) *AddingUser {
	return &AddingUser{service: service}
}

type reqAddUser struct {
	Username string `json:"username"`
}

type resAddUser struct {
	ID int64 `json:"id"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func (h *AddingUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// read request
	var req reqAddUser
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	id, err := h.service.AddUser(req.Username)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// write response
	resp := resAddUser{
		ID: id,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}
