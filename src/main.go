package main

import (
	"achievements/src/controllers"
	"achievements/src/models"
	"achievements/src/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	err := models.ConnectDb()
	if err != nil {
		panic("Failed to connect to database!")
	}

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	r.Route("/challenges", func(r chi.Router) {
		c := controllers.ChallengesWrapper(repository.NewChallengeRepository())
		r.Post("/", c.Create)
		r.Get("/", c.Index)
		r.Get("/{id}", c.Show)
		r.Delete("/{id}", c.Destroy)
	})

	http.ListenAndServe(":3000", r)
}
