package main

import (
	"achievements/src/controllers"
	"achievements/src/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	err := models.ConnectDb()
	if err != nil {
		panic("Failed to connect to database!")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	r.Route("/challenges", func(r chi.Router) {
		c := controllers.Challenges{}
		r.Post("/", c.Create)
	})

	http.ListenAndServe(":3000", r)
}
