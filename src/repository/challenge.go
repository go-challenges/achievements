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
	result := models.DB.Create(challenge)
	if result.Error != nil {
		return result.Error
	}

	user := models.CurrentUser()

	userChallenge := models.UserChallenge{
		UserID:      user.ID,
		ChallengeID: challenge.ID,
	}

	errUserChallenge := userChallenge.Create()
	if errUserChallenge != nil {
		return errUserChallenge
	}

	return nil
}
