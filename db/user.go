package db

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (u *User) AddUser(username string) (int64, error) {
	res, err := u.db.Exec("INSERT INTO users (username, created_at) VALUES (?, ?)", username, time.Now())
	if err != nil {
		if me, ok := err.(*mysql.MySQLError); ok && me.Number == 1062 {
			return 0, fmt.Errorf("user with name '%s' already exists in a db", username)
		}
		return 0, fmt.Errorf("add user to db, %v", err)
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("getting id when added user to db, %v", err)
	}

	return userID, nil
}

func (u *User) CheckExistence(userIDs []int64) error {
	// checking users existence
	query, args, err := sqlx.In("SELECT COUNT(*) FROM users WHERE id IN(?)", userIDs)
	if err != nil {
		return fmt.Errorf("building user check query, %v", err)
	}

	usersInDb := 0
	err = u.db.Get(&usersInDb, query, args...)
	if err != nil {
		return fmt.Errorf("getting count of users from db, %v", err)
	}

	if usersInDb != len(userIDs) {
		return fmt.Errorf("some of users does not exist in db")
	}

	return nil
}
