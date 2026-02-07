// API client for communicating with the Go backend
const API = {
    async getChapters() {
        const res = await fetch('/api/chapters');
        if (!res.ok) throw new Error('Failed to fetch chapters');
        return res.json();
    },

    async getLesson(id) {
        const res = await fetch(`/api/lessons/${id}`);
        if (!res.ok) throw new Error(`Failed to fetch lesson ${id}`);
        return res.json();
    },

    async getQuiz(lessonId) {
        const res = await fetch(`/api/quiz/${lessonId}`);
        if (!res.ok) throw new Error(`Failed to fetch quiz for ${lessonId}`);
        return res.json();
    }
};
