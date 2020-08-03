package handlers

import (
	"encoding/json"
	"messenger/service"
	"net/http"
)

type AddingChat struct {
	service *service.ChatService
}

func NewAddingChat(service *service.ChatService) *AddingChat {
	return &AddingChat{service: service}
}

type reqAddChat struct {
	Name  string  `json:"name"`
	Users []int64 `json:"users"`
}

type resAddChat struct {
	ID int64 `json:"id"`
}

func (h *AddingChat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// read request
	var req reqAddChat
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	id, err := h.service.AddChat(req.Name, req.Users)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// write response
	resp := resAddChat{
		ID: id,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}
