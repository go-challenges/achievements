package main

import (
    "fmt"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    // TODO: add config
    m, err := migrate.New(
        "file://db/migrations",
        "postgres://docker:docker@db:5432/docker?sslmode=disable")

    if err != nil {
		fmt.Println(err);
        return
    }

    if err := m.Up(); err != nil {
		fmt.Println(err);
        return
	}
}