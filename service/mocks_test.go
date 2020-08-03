package service

import (
	"github.com/stretchr/testify/mock"
	"messenger/entities"
)

type userDBMock struct {
	mock.Mock
}

func (m *userDBMock) CheckExistence(userIDs []int64) error {
	args := m.Called(userIDs)
	return args.Error(0)
}

func (m *userDBMock) AddUser(username string) (int64, error) {
	args := m.Called(username)
	return int64(args.Int(0)), args.Error(1)
}

type chatDBMock struct {
	mock.Mock
}

func (m *chatDBMock) AddChat(name string, userIDs []int64) (int64, error) {
	args := m.Called(name, userIDs)
	return int64(args.Int(0)), args.Error(1)
}

func (m *chatDBMock) GetUserChats(userID int64) ([]entities.Chat, error) {
	args := m.Called(userID)
	chats, _ := args.Get(0).([]entities.Chat)
	return chats, args.Error(1)
}

func (m *chatDBMock) UserInChat(chatID int64, userID int64) error {
	args := m.Called(chatID, userID)
	return args.Error(0)
}

type messageDBMock struct {
	mock.Mock
}

func (m *messageDBMock) AddMessage(chatID int64, userID int64, text string) (int64, error) {
	args := m.Called(chatID, userID, text)
	return int64(args.Int(0)), args.Error(1)
}

func (m *messageDBMock) GetChatMessages(chatID int64) ([]entities.Message, error) {
	args := m.Called(chatID)
	messages, _ := args.Get(0).([]entities.Message)
	return messages, args.Error(1)
}
