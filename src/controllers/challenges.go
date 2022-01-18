package controllers

import (
	"achievements/src/forms"
	"achievements/src/models"
	"net/http"

	"github.com/jinzhu/copier"
)

type Challenges struct{}

func (c *Challenges) Create(w http.ResponseWriter, r *http.Request) {
	form := forms.Challenge{}
	if err := form.Bind(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//chal := models.Challenge{Name: form.Name, Description: form.Description}
	chal := models.Challenge{}
	copier.Copy(&chal, &form)

	if result := models.DB.Create(&chal); result.Error != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	Render(w, 201, chal)
}
