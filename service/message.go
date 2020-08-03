package service

import (
	"fmt"
	"messenger/entities"
)

type MessageService struct {
	messageDB MessageDB
	chatDB    ChatDB
}

func NewMessageService(messageDB MessageDB, chatDB ChatDB) *MessageService {
	return &MessageService{messageDB: messageDB, chatDB: chatDB}
}

func (ms *MessageService) AddMessage(chatID int64, userID int64, text string) (int64, error) {
	if err := ms.chatDB.UserInChat(chatID, userID); err != nil {
		return 0, fmt.Errorf("attempt to find user in chat, %v", err)
	}

	id, err := ms.messageDB.AddMessage(chatID, userID, text)
	if err != nil {
		return 0, fmt.Errorf("adding message, %v", err)
	}
	return id, nil
}

func (ms *MessageService) GetChatMessages(chatID int64) ([]entities.Message, error) {
	messages, err := ms.messageDB.GetChatMessages(chatID)
	if err != nil {
		return nil, fmt.Errorf("getting all chat's messages, %v", err)
	}
	return messages, nil
}
