package models

import (
	"time"
)

type User struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	// TODO: add deleted at
}

func CurrentUser() User {
	var user User
	DB.Last(&user)

	return user
}
