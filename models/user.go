package models

import "time"

type User struct {
	ID             int
	Name           string
	Email          string
	Role           string
	Username       string
	Password       string
	FailedAttempts int
	PasswordExpiry time.Time
	TelegramID     string
	SecretKey      string
	IsConfirm      bool
	Remarks        string
}
