package entities

import "time"

type Chat struct {
	ID              int64
	Name            string
	CreatedAt       time.Time  `db:"created_at"`
	LastMessageTime *time.Time `db:"last_message_time"`
}
