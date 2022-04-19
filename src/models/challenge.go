package models

import (
	"time"

	"achievements/src/helpers"

	"gorm.io/gorm"
)

type Challenge struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Status      string         `gorm:"type:challenge_statuses;default:created;" json:"status"`
	Type        string         `gorm:"type:challenge_types;default:number;" json:"type"`
	Properties  helpers.JSONB  `gorm:"type:jsonb" json:"properties,omitempty"`
}

type ChallengeRepository interface {
	Create() error
}

func (c *Challenge) Create() error {
	err := DB.Create(c)
	if err != nil {
		return err.Error
	}

	return nil
}
