package entities_user_v1

import "time"

type User struct {
	ID             string
	Username       string
	IsVerified     bool
	ProfilePicture string
	CreatedAt      time.Time
}
