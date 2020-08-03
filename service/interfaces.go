package service

import "messenger/entities"

type UserDB interface {
	AddUser(username string) (int64, error)
	CheckExistence(userIDs []int64) error
}

type ChatDB interface {
	AddChat(name string, userIDs []int64) (int64, error)
	GetUserChats(userID int64) ([]entities.Chat, error)
	UserInChat(chatID int64, userID int64) error
}

type MessageDB interface {
	AddMessage(chatID int64, userID int64, text string) (int64, error)
	GetChatMessages(chatID int64) ([]entities.Message, error)
}
