package controllers

import (
	"achievements/src/forms"
	"achievements/src/models"
	"achievements/src/repository"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jinzhu/copier"
	"gorm.io/gorm/clause"
)

type Challenges struct {
	repository repository.ChallengesRepository
}

func ChallengesWrapper(ChallRepo repository.ChallengesRepository) *Challenges {
	return &Challenges{
		repository: ChallRepo,
	}
}

func (c *Challenges) Create(w http.ResponseWriter, r *http.Request) {
	form := forms.Challenge{}
	if err := form.Bind(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//chal := models.Challenge{Name: form.Name, Description: form.Description}
	challenge := models.Challenge{}
	copier.Copy(&challenge, &form)

	if err := c.repository.Create(&challenge); err != nil {
		fmt.Println("Shit for db")
		http.Error(w, http.StatusText(422), 422)
		return
	}

	Render(w, 201, challenge)
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

/*
Актор пользователь
1. Пользователь заходит на сайт
2. Пользователь видит список challanges (общий)
3. У пользователя есть два пути - присоединиться и создать
4.

Модель challange
id
created_at
updated_at
delete_at
name
description
state (in_moderation, active, ended)
type
rules (????)
schedule (???)
permissions (????) - вторая фаза

state (pending, declined - as invitation maybe???) , completed, in_progress, failed


Типы
- числовой
- числовой  + время
- олимпийка
- сетка
- захват флага

ChallengeProperties (относиться к типу только)
общие правила для всех пользователй
определяет как вносить данные
????

Schedule
???

список challaeges
категории чтобы поиск упростить

Создание challenge
1. Пользователь нажимает на кнопку создать
2. Видит форму - название, описание, под-форма календаря (начало, конец, период напоминания), под-форма правил, тип
3. Пользователь заполняет форму
4. Нажимает save

Присоединение в существующий challenge
Автоматическое или по инвайту владельца


Создание Participation(?) (в зависимоти от типа)
Для конкретного юзера отметка что он сделал
Мануальное или автоматическое

Закрытие challenge




3, - активные в которых он учавствует
4. Пользователь видит кнопку добавить challenge
5.
*/
