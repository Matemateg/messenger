package entities

import "time"

type Message struct {
	ID        int64
	ChatID    int64 `db:"chat_id"`
	UserID    int64 `db:"author_id"`
	Text      string
	CreatedAt time.Time `db:"created_at"`
}
