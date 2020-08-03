package handlers

import (
	"encoding/json"
	"messenger/service"
	"net/http"
)

type AddingMessage struct {
	service *service.MessageService
}

func NewAddingMessage(service *service.MessageService) *AddingMessage {
	return &AddingMessage{service: service}
}

type reqAddMessage struct {
	ChatID int64  `json:"chat"`
	UserID int64  `json:"author"`
	Text   string `json:"text"`
}

type resAddMessage struct {
	ID int64 `json:"id"`
}

func (h *AddingMessage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// read request
	var req reqAddMessage
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	id, err := h.service.AddMessage(req.ChatID, req.UserID, req.Text)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// write response
	resp := resAddMessage{
		ID: id,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}
