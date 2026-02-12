package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"go-learning-app/data"
	"go-learning-app/handlers"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	// Initialize SQLite database
	db, err := data.NewDB(data.DBPath())
	if err != nil {
		log.Fatalf("データベースの初期化に失敗しました: %v", err)
	}
	defer db.Close()

	store := data.NewStore()
	h := handlers.New(store, db)

	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("GET /api/chapters", h.GetChapters)
	mux.HandleFunc("GET /api/lessons/{id}", h.GetLesson)
	mux.HandleFunc("GET /api/quiz/{lessonId}", h.GetQuiz)
	mux.HandleFunc("POST /api/run", h.RunCode)

	// Progress API routes
	mux.HandleFunc("POST /api/login", h.Login)
	mux.HandleFunc("GET /api/progress/{username}", h.GetProgress)
	mux.HandleFunc("POST /api/progress/{username}/{lessonId}", h.MarkProgress)
	mux.HandleFunc("DELETE /api/progress/{username}", h.ResetProgress)

	// Static files
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("GET /", http.FileServer(http.FS(staticFS)))

	// Get port from environment variable (Cloud Run sets $PORT)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	fmt.Printf("Go学習アプリを起動しました: http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
