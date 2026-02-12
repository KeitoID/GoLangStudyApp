package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"go-learning-app/data"
)

// Handler holds the data store and provides HTTP handler methods.
type Handler struct {
	store *data.Store
	db    *data.DB
}

// New creates a new Handler with the given store and database.
func New(store *data.Store, db *data.DB) *Handler {
	return &Handler{store: store, db: db}
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

// Login handles user login/registration and returns progress.
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request"})
		return
	}

	username := strings.TrimSpace(req.Username)
	if username == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "ユーザー名を入力してください"})
		return
	}

	if _, err := h.db.EnsureUser(username); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to create user"})
		return
	}

	progress, err := h.db.GetProgress(username)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to get progress"})
		return
	}
	if progress == nil {
		progress = []string{}
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"username": username,
		"progress": progress,
	})
}

// GetProgress returns completed lessons for a user.
func (h *Handler) GetProgress(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	progress, err := h.db.GetProgress(username)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to get progress"})
		return
	}
	if progress == nil {
		progress = []string{}
	}

	writeJSON(w, http.StatusOK, map[string]any{"progress": progress})
}

// MarkProgress marks a lesson as completed for a user.
func (h *Handler) MarkProgress(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	lessonID := r.PathValue("lessonId")

	if err := h.db.MarkCompleted(username, lessonID); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to save progress"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

// ResetProgress deletes all progress for a user.
func (h *Handler) ResetProgress(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	if err := h.db.ResetProgress(username); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to reset progress"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

// RunCodeRequest is the request body for running code.
type RunCodeRequest struct {
	Code string `json:"code"`
}

// RunCodeResponse is the response from running code.
type RunCodeResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

// RunCode executes Go code and returns the output.
func (h *Handler) RunCode(w http.ResponseWriter, r *http.Request) {
	var req RunCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, RunCodeResponse{Error: "Invalid request"})
		return
	}

	if req.Code == "" {
		writeJSON(w, http.StatusBadRequest, RunCodeResponse{Error: "コードが空です"})
		return
	}

	// Create temp directory for code execution
	tmpDir, err := os.MkdirTemp("", "gorun-*")
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, RunCodeResponse{Error: "Failed to create temp directory"})
		return
	}
	defer os.RemoveAll(tmpDir)

	// Write code to temp file
	codePath := filepath.Join(tmpDir, "main.go")
	if err := os.WriteFile(codePath, []byte(req.Code), 0644); err != nil {
		writeJSON(w, http.StatusInternalServerError, RunCodeResponse{Error: "Failed to write code"})
		return
	}

	// Execute with timeout (5 seconds)
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", codePath)
	output, err := cmd.CombinedOutput()

	resp := RunCodeResponse{}
	if ctx.Err() == context.DeadlineExceeded {
		resp.Error = "実行がタイムアウトしました（5秒）"
		resp.Output = string(output)
	} else if err != nil {
		resp.Error = "実行エラー"
		resp.Output = string(output)
	} else {
		resp.Output = string(output)
	}

	writeJSON(w, http.StatusOK, resp)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
