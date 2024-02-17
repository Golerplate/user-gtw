package entities_session_v1

import "time"

type Session struct {
	ID        string
	UserID    string
	CreatedAt time.Time
}
