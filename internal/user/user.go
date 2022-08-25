package user

import (
	"context"
)

type User struct {
	UserID       int
	Username     string
	ActiveTicker bool

	Ctx context.Context
}

func CreateUser(ctx context.Context, username string, userID int) *User {
	return &User{
		UserID:       userID,
		Username:     username,
		ActiveTicker: false,
		Ctx:          ctx,
	}
}
