package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"messenger/entities"
	"time"
)

type Message struct {
	db *sqlx.DB
}

func NewMessage(db *sqlx.DB) *Message {
	return &Message{db: db}
}

func (m *Message) AddMessage(chatID int64, userID int64, text string) (int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("begin transaction, %v", err)
	}
	res, err := m.db.Exec("INSERT INTO messages (chat_id, author_id, text, created_at) VALUES (?, ?, ?, ?)", chatID, userID, text, time.Now())
	if err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("add user to db, %v", err)
	}

	messageID, err := res.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("getting id when added message to db, %v", err)
	}

	_, err = m.db.Exec(
		`UPDATE chats 
		SET last_message_time = (SELECT created_at FROM messages WHERE id = ?)
		WHERE id = ?`,
		messageID,
		chatID,
	)
	if err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("update last message time, %v", err)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("commit, %v", err)
	}

	return messageID, nil
}

func (m *Message) GetChatMessages(chatID int64) ([]entities.Message, error) {
	rows, err := m.db.Queryx(
		`SELECT * FROM messages WHERE chat_id = ? ORDER BY created_at ASC`, chatID)
	if err != nil {
		return nil, err
	}

	var messages []entities.Message
	for rows.Next() {
		var message entities.Message
		if err = rows.StructScan(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
