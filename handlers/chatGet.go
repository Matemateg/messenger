package handlers

import (
	"encoding/json"
	"messenger/service"
	"net/http"
	"time"
)

type GettingUserChats struct {
	service *service.ChatService
}

func NewGettingUserChats(service *service.ChatService) *GettingUserChats {
	return &GettingUserChats{service: service}
}

type reqGetChat struct {
	UserID int64 `json:"user"`
}

type resGetChat struct {
	ID              int64      `json:"id"`
	Name            string     `json:"name"`
	CreatedAt       time.Time  `json:"created_at"`
	LastMessageTime *time.Time `json:"last_message_time"`
}

func (h *GettingUserChats) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// read request
	var req reqGetChat
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	chats, err := h.service.GetUserChats(req.UserID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	var resp []resGetChat

	for _, chat := range chats {
		resp = append(resp, resGetChat(chat))
	}

	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}
