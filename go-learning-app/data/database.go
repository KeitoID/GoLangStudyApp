package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

// DB wraps an SQLite connection for progress persistence.
type DB struct {
	conn *sql.DB
}

// NewDB opens the SQLite database at path and creates tables if needed.
func NewDB(path string) (*DB, error) {
	conn, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	// Enable WAL mode for better concurrent read performance.
	if _, err := conn.Exec("PRAGMA journal_mode=WAL"); err != nil {
		conn.Close()
		return nil, fmt.Errorf("set wal mode: %w", err)
	}

	if err := createTables(conn); err != nil {
		conn.Close()
		return nil, err
	}

	return &DB{conn: conn}, nil
}

func createTables(conn *sql.DB) error {
	const schema = `
CREATE TABLE IF NOT EXISTS users (
    username   TEXT PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS progress (
    username     TEXT NOT NULL,
    lesson_id    TEXT NOT NULL,
    completed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (username, lesson_id),
    FOREIGN KEY (username) REFERENCES users(username)
);`

	if _, err := conn.Exec(schema); err != nil {
		return fmt.Errorf("create tables: %w", err)
	}
	return nil
}

// Close closes the database connection.
func (db *DB) Close() error {
	return db.conn.Close()
}

// EnsureUser creates the user if it doesn't exist. Returns true if newly created.
func (db *DB) EnsureUser(username string) (bool, error) {
	res, err := db.conn.Exec(
		"INSERT OR IGNORE INTO users (username) VALUES (?)", username,
	)
	if err != nil {
		return false, fmt.Errorf("ensure user: %w", err)
	}
	n, _ := res.RowsAffected()
	return n > 0, nil
}

// GetProgress returns the list of completed lesson IDs for a user.
func (db *DB) GetProgress(username string) ([]string, error) {
	rows, err := db.conn.Query(
		"SELECT lesson_id FROM progress WHERE username = ? ORDER BY completed_at",
		username,
	)
	if err != nil {
		return nil, fmt.Errorf("get progress: %w", err)
	}
	defer rows.Close()

	var lessons []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan progress: %w", err)
		}
		lessons = append(lessons, id)
	}
	return lessons, rows.Err()
}

// MarkCompleted records that the user completed a lesson.
func (db *DB) MarkCompleted(username, lessonID string) error {
	_, err := db.conn.Exec(
		"INSERT OR IGNORE INTO progress (username, lesson_id) VALUES (?, ?)",
		username, lessonID,
	)
	if err != nil {
		return fmt.Errorf("mark completed: %w", err)
	}
	return nil
}

// ResetProgress deletes all progress records for a user.
func (db *DB) ResetProgress(username string) error {
	_, err := db.conn.Exec("DELETE FROM progress WHERE username = ?", username)
	if err != nil {
		return fmt.Errorf("reset progress: %w", err)
	}
	return nil
}

// DBPath returns the database file path from env or default.
func DBPath() string {
	if p := os.Getenv("DB_PATH"); p != "" {
		return p
	}
	return "go-learning.db"
}
