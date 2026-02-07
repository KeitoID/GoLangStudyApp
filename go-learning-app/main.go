package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"go-learning-app/data"
	"go-learning-app/handlers"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	store := data.NewStore()
	h := handlers.New(store)

	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("GET /api/chapters", h.GetChapters)
	mux.HandleFunc("GET /api/lessons/{id}", h.GetLesson)
	mux.HandleFunc("GET /api/quiz/{lessonId}", h.GetQuiz)

	// Static files
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("GET /", http.FileServer(http.FS(staticFS)))

	addr := ":8080"
	fmt.Printf("Go学習アプリを起動しました: http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
