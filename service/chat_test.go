package service

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"messenger/entities"
	"testing"
)

func TestChatService_AddChat(t *testing.T) {
	t.Run("creating chat without users", func(t *testing.T) {
		chatSrv := NewChatService(nil, nil)
		_, err := chatSrv.AddChat("cats", nil)
		require.EqualError(t, err, "can't create chat without users")
	})

	t.Run("creating chat with empty chatname returns an error", func(t *testing.T) {
		userSrv := NewChatService(nil, nil)
		_, err := userSrv.AddChat("", []int64{1})
		require.EqualError(t, err, "chatname is too short")
	})

	t.Run("creating chat with nonexistent users returns an error", func(t *testing.T) {
		dbChat := &chatDBMock{}
		dbUser := &userDBMock{}
		dbUser.On("CheckExistence", mock.Anything).Return(errors.New("users does not exist"))

		chatSrv := NewChatService(dbChat, dbUser)
		_, err := chatSrv.AddChat("cats", []int64{1})
		require.EqualError(t, err, "users not exist in db, users does not exist")
	})

	t.Run("creating chat with error in db.AddChat returns an error", func(t *testing.T) {
		dbChat := &chatDBMock{}
		dbUser := &userDBMock{}
		dbUser.On("CheckExistence", mock.Anything).Return(nil)
		dbChat.On("AddChat", mock.Anything, mock.Anything).Return(0, errors.New("a db error"))

		chatSrv := NewChatService(dbChat, dbUser)
		_, err := chatSrv.AddChat("cats", []int64{1})
		require.EqualError(t, err, "adding chat, a db error")
	})

	t.Run("creating chat returns id", func(t *testing.T) {
		dbChat := &chatDBMock{}
		dbUser := &userDBMock{}
		dbUser.On("CheckExistence", mock.Anything).Return(nil)
		dbChat.On("AddChat", mock.Anything, mock.Anything).Return(11, nil)

		chatSrv := NewChatService(dbChat, dbUser)
		id, err := chatSrv.AddChat("cats", []int64{1})
		require.NoError(t, err)
		require.Equal(t, int64(11), id)
	})
}

func TestChatService_GetUserChats(t *testing.T) {
	t.Run("returns error on error db", func(t *testing.T) {
		dbChat := &chatDBMock{}
		dbChat.On("GetUserChats", mock.Anything).Return(nil, errors.New("db error"))

		chatSrv := NewChatService(dbChat, nil)
		_, err := chatSrv.GetUserChats(1)
		require.EqualError(t, err, "getting all user's chats, db error")
	})

	t.Run("returns chat list", func(t *testing.T) {
		dbChat := &chatDBMock{}
		expectedResult := []entities.Chat{
			{ID: 1, Name: "cats"},
			{ID: 2, Name: "dogs"},
		}
		dbChat.On("GetUserChats", mock.Anything).Return(expectedResult, nil)

		chatSrv := NewChatService(dbChat, nil)
		chatList, err := chatSrv.GetUserChats(1)
		require.NoError(t, err)
		require.Equal(t, expectedResult, chatList)
	})
}
