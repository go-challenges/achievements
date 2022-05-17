package repository

import (
	"achievements/src/helpers"
	"achievements/src/models"
	"time"
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

	duration, err := helpers.FindPeriod(challenge.Properties["duration"])
	if err != nil {
		return err
	}
	time := time.Now()
	userChallenge := models.UserChallenge{
		UserID:      user.ID,
		ChallengeID: challenge.ID,
		StartedAt:   &time,
		EndedAt:     &duration,
	}

	errUserChallenge := userChallenge.Create()
	if errUserChallenge != nil {
		return errUserChallenge
	}

	return nil
}
