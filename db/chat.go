package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"messenger/entities"
	"time"
)

type Chat struct {
	db *sqlx.DB
}

func NewChat(db *sqlx.DB) *Chat {
	return &Chat{db: db}
}

func (ch *Chat) AddChat(name string, userIDs []int64) (int64, error) {
	tx, err := ch.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("begin transaction, %v", err)
	}
	res, err := tx.Exec("INSERT INTO chats (name, created_at) VALUES (?, ?)", name, time.Now())
	if err != nil {
		_ = tx.Rollback()
		if me, ok := err.(*mysql.MySQLError); ok && me.Number == 1062 {
			return 0, fmt.Errorf("chat with name '%s' already exists in db", name)
		}
		return 0, fmt.Errorf("add chat to db, %v", err)
	}

	chatID, err := res.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("getting id when added chat to db, %v", err)
	}

	for _, user := range userIDs {
		_, err = tx.Exec("INSERT INTO chats_has_users (chat_id, user_id) VALUES (?, ?)", chatID, user)
		if err != nil {
			_ = tx.Rollback()
			return 0, fmt.Errorf("add chat_has_user to db, %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("commit, %v", err)
	}

	return chatID, nil
}

func (ch *Chat) GetUserChats(userID int64) ([]entities.Chat, error) {
	rows, err := ch.db.Queryx(
		`SELECT c.*
				FROM chats c
				JOIN chats_has_users chu
				ON c.id = chu.chat_id
				WHERE user_id = ?
				ORDER BY c.last_message_time DESC`,
		userID)
	if err != nil {
		return nil, err
	}

	var chats []entities.Chat
	for rows.Next() {
		var chat entities.Chat
		if err = rows.StructScan(&chat); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	return chats, nil
}

func (ch *Chat) UserInChat(chatID int64, userID int64) error {
	//checking if the user is in the chat
	row := ch.db.QueryRowx("SELECT 1 FROM chats_has_users WHERE chat_id = ? AND user_id = ?", chatID, userID)
	var dummy int8
	err := row.Scan(&dummy)
	if err == sql.ErrNoRows {
		return fmt.Errorf("user isn't in chat")
	}
	if err != nil {
		return fmt.Errorf("getting pair chat-user, %v", err)
	}
	return nil
}
