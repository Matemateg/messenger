package service

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserService_AddUser(t *testing.T) {
	t.Run("creating user with empty username returns an error", func(t *testing.T) {
		userSrv := NewUserService(nil)
		_, err := userSrv.AddUser("")
		require.EqualError(t, err, "username is too short")
	})

	t.Run("creating user returns error on error db", func(t *testing.T) {
		db := &userDBMock{}
		db.On("AddUser", "vova").Return(0, errors.New("db failed"))

		userSrv := NewUserService(db)
		_, err := userSrv.AddUser("vova")
		require.EqualError(t, err, "adding user, db failed")
	})

	t.Run("creating user returns id", func(t *testing.T) {
		db := &userDBMock{}
		db.On("AddUser", "vova").Return(11, nil)

		userSrv := NewUserService(db)
		id, err := userSrv.AddUser("vova")
		require.NoError(t, err)
		require.Equal(t, int64(11), id)
	})
}
