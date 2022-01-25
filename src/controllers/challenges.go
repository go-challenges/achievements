package controllers

import (
	"achievements/src/forms"
	"achievements/src/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jinzhu/copier"
	"gorm.io/gorm/clause"
)

type Challenges struct{}

func (c *Challenges) Create(w http.ResponseWriter, r *http.Request) {
	form := forms.Challenge{}
	if err := form.Bind(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//chal := models.Challenge{Name: form.Name, Description: form.Description}
	challange := models.Challenge{}
	copier.Copy(&challange, &form)

	if result := models.DB.Create(&challange); result.Error != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	Render(w, 201, challange)
}

func (c *Challenges) Index(w http.ResponseWriter, r *http.Request) {
	challenges := []models.Challenge{}
	models.DB.Order("created_at").Find(&challenges)

	if len(challenges) == 0 {
		RenderEmpty(w, 404)
		return
	}

	Render(w, 200, challenges)
}

func (c *Challenges) Show(w http.ResponseWriter, r *http.Request) {
	challenge := models.Challenge{}
	models.DB.Find(&challenge, chi.URLParam(r, "id"))

	if challenge.ID == 0 {
		RenderEmpty(w, 404)
		return
	}

	Render(w, 200, challenge)
}

func (c *Challenges) Destroy(w http.ResponseWriter, r *http.Request) {
	challenge := models.Challenge{}
	models.DB.Clauses(clause.Returning{}).Delete(&challenge, chi.URLParam(r, "id"))

	if challenge.ID == 0 {
		RenderEmpty(w, 404)
		return
	}

	RenderEmpty(w, 204)
}
