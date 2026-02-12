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
    },

    async login(username) {
        const res = await fetch('/api/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username }),
        });
        if (!res.ok) {
            const data = await res.json();
            throw new Error(data.error || 'Login failed');
        }
        return res.json();
    },

    async getProgress(username) {
        const res = await fetch(`/api/progress/${encodeURIComponent(username)}`);
        if (!res.ok) throw new Error('Failed to fetch progress');
        return res.json();
    },

    async markCompleted(username, lessonId) {
        const res = await fetch(
            `/api/progress/${encodeURIComponent(username)}/${encodeURIComponent(lessonId)}`,
            { method: 'POST' },
        );
        if (!res.ok) throw new Error('Failed to save progress');
        return res.json();
    },

    async resetProgress(username) {
        const res = await fetch(`/api/progress/${encodeURIComponent(username)}`, {
            method: 'DELETE',
        });
        if (!res.ok) throw new Error('Failed to reset progress');
        return res.json();
    },
};
