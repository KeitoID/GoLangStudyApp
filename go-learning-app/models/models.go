package models

// Chapter represents a learning chapter containing multiple lessons.
type Chapter struct {
	ID          int             `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Lessons     []LessonSummary `json:"lessons"`
}

// LessonSummary is a brief view of a lesson used in chapter listings.
type LessonSummary struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// Lesson is the full lesson content including code examples and notes.
type Lesson struct {
	ID          string        `json:"id"`
	ChapterID   int           `json:"chapterId"`
	Title       string        `json:"title"`
	Content     string        `json:"content"`
	CodeExamples []CodeExample `json:"codeExamples"`
	Notes       []string      `json:"notes,omitempty"`
}

// CodeExample holds a titled code snippet.
type CodeExample struct {
	Title string `json:"title"`
	Code  string `json:"code"`
}

// Quiz holds the questions for a particular lesson.
type Quiz struct {
	LessonID  string     `json:"lessonId"`
	Questions []Question `json:"questions"`
}

// Question is a single multiple-choice quiz question.
type Question struct {
	ID          string   `json:"id"`
	Text        string   `json:"text"`
	Options     []string `json:"options"`
	Answer      int      `json:"answer"`
	Explanation string   `json:"explanation"`
}
