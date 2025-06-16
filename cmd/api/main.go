package main

import (
	"database/sql"
	"event/internal/database"
	"event/internal/env"
	"log"
	_"github.com/mattn/go-sqlite3"
	_"github.com/joho/godotenv/autoload"
)
type application struct {
	port int
	jwtSecret string
	models database.Models
}
func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)
	app := &application {
		port: env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-secret-here" ),
		models: models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}

}