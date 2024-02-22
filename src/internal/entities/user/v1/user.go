package entities_user_v1

import "time"

type User struct {
	ID         string
	ExternalID string
	Username   string
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
