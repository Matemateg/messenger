package service

import (
	"fmt"
)

type UserService struct {
	userDB UserDB
}

func NewUserService(userDB UserDB) *UserService {
	return &UserService{userDB: userDB}
}

func (us *UserService) AddUser(username string) (int64, error) {
	if username == "" {
		return 0, fmt.Errorf("username is too short")
	}
	id, err := us.userDB.AddUser(username)
	if err != nil {
		return 0, fmt.Errorf("adding user, %v", err)
	}
	return id, nil
}
