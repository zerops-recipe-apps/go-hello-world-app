package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/", handleRoot)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type healthResponse struct {
	Type     string            `json:"type"`
	Greeting string            `json:"greeting"`
	Status   map[string]string `json:"status"`
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	resp := healthResponse{
		Type:   "go",
		Status: make(map[string]string),
	}

	if err := db.Ping(); err != nil {
		resp.Status["database"] = fmt.Sprintf("ERROR: %v", err)
		writeJSON(w, http.StatusServiceUnavailable, resp)
		return
	}

	if err := db.QueryRow("SELECT message FROM greetings LIMIT 1").Scan(&resp.Greeting); err != nil {
		resp.Status["database"] = fmt.Sprintf("ERROR: %v", err)
		writeJSON(w, http.StatusServiceUnavailable, resp)
		return
	}

	resp.Status["database"] = "OK"
	writeJSON(w, http.StatusOK, resp)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}
