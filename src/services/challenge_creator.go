package services

import (
	"achievements/src/forms"
	"achievements/src/models"
	"achievements/src/repository"

	"github.com/jinzhu/copier"
)

func CreateChallenge(form *forms.Challenge, repository repository.ChallengesRepository) (*models.Challenge, error) {
	challenge := models.Challenge{}
	copier.Copy(&challenge, &form)

	if err := repository.Create(&challenge); err != nil {
		return nil, err
	}

	return &challenge, nil
}
