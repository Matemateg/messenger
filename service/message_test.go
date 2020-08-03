package service

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"messenger/entities"
	"testing"
)

func TestMessageService_AddMessage(t *testing.T) {
	t.Run("returns an error if user is not in the chat", func(t *testing.T) {
		dbChat := &chatDBMock{}
		dbMessage := &messageDBMock{}
		dbChat.On("UserInChat", mock.Anything, mock.Anything).Return(errors.New("user is not in the chat"))

		msgSrv := NewMessageService(dbMessage, dbChat)
		_, err := msgSrv.AddMessage(1, 3, "222")
		require.EqualError(t, err, "attempt to find user in chat, user is not in the chat")
	})

	t.Run("returns an error if can't add message to db", func(t *testing.T) {
		dbChat := &chatDBMock{}
		dbMessage := &messageDBMock{}
		dbChat.On("UserInChat", mock.Anything, mock.Anything).Return(nil)
		dbMessage.On("AddMessage", mock.Anything, mock.Anything, mock.Anything).Return(0, errors.New("an error"))

		msgSrv := NewMessageService(dbMessage, dbChat)
		_, err := msgSrv.AddMessage(1, 3, "222")
		require.EqualError(t, err, "adding message, an error")
	})

	t.Run("returns message id", func(t *testing.T) {
		dbChat := &chatDBMock{}
		dbMessage := &messageDBMock{}
		dbChat.On("UserInChat", mock.Anything, mock.Anything).Return(nil)
		dbMessage.On("AddMessage", mock.Anything, mock.Anything, mock.Anything).Return(22, nil)

		msgSrv := NewMessageService(dbMessage, dbChat)
		id, err := msgSrv.AddMessage(1, 3, "222")
		require.NoError(t, err)
		require.Equal(t, int64(22), id)
	})
}

func TestMessageService_GetChatMessages(t *testing.T) {
	t.Run("returns en error if db returns an error", func(t *testing.T) {
		dbChat := &chatDBMock{}
		dbMessage := &messageDBMock{}
		dbMessage.On("GetChatMessages", mock.Anything).Return(nil, errors.New("an error"))
		msgSrv := NewMessageService(dbMessage, dbChat)
		_, err := msgSrv.GetChatMessages(1)
		require.EqualError(t, err, "getting all chat's messages, an error")
	})

	t.Run("returns messages list", func(t *testing.T) {
		dbChat := &chatDBMock{}
		dbMessage := &messageDBMock{}
		expectedMsgList := []entities.Message{
			{ID: 1, Text: "Hi volodya"},
			{ID: 2, Text: "What's up boris?"},
		}
		dbMessage.On("GetChatMessages", mock.Anything).Return(expectedMsgList, nil)
		msgSrv := NewMessageService(dbMessage, dbChat)
		msgList, err := msgSrv.GetChatMessages(1)
		require.NoError(t, err)
		require.Equal(t, expectedMsgList, msgList)
	})
}
