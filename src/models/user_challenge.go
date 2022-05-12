package models

import (
	"time"
)

type UserChallenge struct {
	ID          int        `gorm:"primaryKey" json:"id"`
	UserID      int        `json:"user_id"`
	ChallengeID int        `json:"challenge_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type Tabler interface {
	TableName() string
}

func (UserChallenge) TableName() string {
	return "users_challenges"
}

func (c *UserChallenge) Create() error {
	err := DB.Create(c)
	if err != nil {
		return err.Error
	}

	return nil
}
