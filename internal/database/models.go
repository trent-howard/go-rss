// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package database

import (
	"time"
)

type Feed struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    int32
}

type User struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}
