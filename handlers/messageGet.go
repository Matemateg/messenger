package handlers

import (
	"encoding/json"
	"messenger/service"
	"net/http"
	"time"
)

type GettingChatMessages struct {
	service *service.MessageService
}

func NewGettingChatMessages(service *service.MessageService) *GettingChatMessages {
	return &GettingChatMessages{service: service}
}

type reqGetMessage struct {
	ChatID int64 `json:"chat"`
}

type resGetMessage struct {
	ID        int64     `json:"id"`
	ChatID    int64     `json:"chat"`
	UserID    int64     `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *GettingChatMessages) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// read request
	var req reqGetMessage
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	messages, err := h.service.GetChatMessages(req.ChatID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	var resp []resGetMessage

	for _, message := range messages {
		resp = append(resp, resGetMessage(message))
	}

	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}
