package service

import (
	"fmt"
	"messenger/entities"
)

type ChatService struct {
	chatDB ChatDB
	userDB UserDB
}

func NewChatService(chatDB ChatDB, userDB UserDB) *ChatService {
	return &ChatService{chatDB: chatDB, userDB: userDB}
}

func (chs *ChatService) AddChat(name string, userIDs []int64) (int64, error) {
	if len(userIDs) == 0 {
		return 0, fmt.Errorf("can't create chat without users")
	}

	if name == "" {
		return 0, fmt.Errorf("chatname is too short")
	}

	if err := chs.userDB.CheckExistence(userIDs); err != nil {
		return 0, fmt.Errorf("users not exist in db, %v", err)
	}

	id, err := chs.chatDB.AddChat(name, userIDs)
	if err != nil {
		return 0, fmt.Errorf("adding chat, %v", err)
	}
	return id, nil
}

func (chs *ChatService) GetUserChats(userID int64) ([]entities.Chat, error) {
	chats, err := chs.chatDB.GetUserChats(userID)
	if err != nil {
		return nil, fmt.Errorf("getting all user's chats, %v", err)
	}
	return chats, nil
}
