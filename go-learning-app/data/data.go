package data

import "go-learning-app/models"

// Store holds all chapters, lessons, and quizzes with indexed lookup.
type Store struct {
	Chapters []models.Chapter
	lessons  map[string]models.Lesson
	quizzes  map[string]models.Quiz
}

// NewStore creates and initializes the data store with all content.
func NewStore() *Store {
	s := &Store{
		lessons: make(map[string]models.Lesson),
		quizzes: make(map[string]models.Quiz),
	}
	s.loadAll()
	return s
}

// GetChapters returns all chapters with lesson summaries.
func (s *Store) GetChapters() []models.Chapter {
	return s.Chapters
}

// GetLesson returns a lesson by ID (e.g. "1-1").
func (s *Store) GetLesson(id string) (models.Lesson, bool) {
	l, ok := s.lessons[id]
	return l, ok
}

// GetQuiz returns a quiz by lesson ID.
func (s *Store) GetQuiz(lessonID string) (models.Quiz, bool) {
	q, ok := s.quizzes[lessonID]
	return q, ok
}

func (s *Store) addChapter(ch models.Chapter) {
	s.Chapters = append(s.Chapters, ch)
}

func (s *Store) addLesson(l models.Lesson) {
	s.lessons[l.ID] = l
}

func (s *Store) addQuiz(q models.Quiz) {
	s.quizzes[q.LessonID] = q
}

func (s *Store) loadAll() {
	loadChapter01(s)
	loadChapter02(s)
	loadChapter03(s)
	loadChapter04(s)
	loadChapter05(s)
	loadChapter06(s)
	loadChapter07(s)
	loadChapter08(s)
	loadChapter09(s)
	loadChapter10(s)
}
