package entity

import "time"

type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash string // hash cünkü raw password business object icinde tutulmaz
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
