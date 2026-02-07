package handlers

import (
	"encoding/json"
	"net/http"

	"go-learning-app/data"
)

// Handler holds the data store and provides HTTP handler methods.
type Handler struct {
	store *data.Store
}

// New creates a new Handler with the given store.
func New(store *data.Store) *Handler {
	return &Handler{store: store}
}

// GetChapters returns all chapters as JSON.
func (h *Handler) GetChapters(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.store.GetChapters())
}

// GetLesson returns a single lesson by ID.
func (h *Handler) GetLesson(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	lesson, ok := h.store.GetLesson(id)
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "lesson not found"})
		return
	}
	writeJSON(w, http.StatusOK, lesson)
}

// GetQuiz returns the quiz for a lesson.
func (h *Handler) GetQuiz(w http.ResponseWriter, r *http.Request) {
	lessonID := r.PathValue("lessonId")
	quiz, ok := h.store.GetQuiz(lessonID)
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "quiz not found"})
		return
	}
	writeJSON(w, http.StatusOK, quiz)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
