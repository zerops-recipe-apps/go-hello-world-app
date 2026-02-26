package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS greetings (
			id      INTEGER PRIMARY KEY,
			message TEXT    NOT NULL
		);
		INSERT INTO greetings (id, message)
		VALUES (1, 'Hello from Zerops!')
		ON CONFLICT (id) DO NOTHING;
	`)
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("migration complete")
}
