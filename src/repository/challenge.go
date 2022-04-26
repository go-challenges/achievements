package repository

import (
	"achievements/src/models"
)

type ChallengeRepository struct{}

type ChallengesRepository interface {
	Create(*models.Challenge) error
}

func NewChallengeRepository() *ChallengeRepository {
	return &ChallengeRepository{}
}

func (r *ChallengeRepository) Create(challenge *models.Challenge) error {
	err := models.DB.Create(challenge)
	if err != nil {
		return err.Error
	}

	return nil
}
