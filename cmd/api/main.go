package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"test_for_goforge/internal/handler"
	"test_for_goforge/internal/repository"
	"time"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// Ждем готовности БД
	for i := 0; i < 30; i++ {
		if err := db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	repo := repository.New(db)
	h := handler.New(repo)

	http.HandleFunc("/number", h.AddNumber)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
