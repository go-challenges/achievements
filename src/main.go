package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Challenge struct {
	gorm.Model
	Name        string
	Description string
}

func main() {
	dsn := "host=db user=docker password=docker dbname=docker port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Shit!")
	}

	chal := Challenge{Name: "First Clng", Description: "First chal descr"}
	result := db.Create(&chal)
	fmt.Println(result)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	http.ListenAndServe(":3000", r)
}
