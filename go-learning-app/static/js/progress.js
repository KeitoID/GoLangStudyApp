// Progress management using server-side SQLite via API
const Progress = {
    USERNAME_KEY: 'go-learning-username',
    username: null,
    _completed: new Set(),

    getUsername() {
        if (this.username) return this.username;
        this.username = localStorage.getItem(this.USERNAME_KEY);
        return this.username;
    },

    setUsername(name) {
        this.username = name;
        localStorage.setItem(this.USERNAME_KEY, name);
    },

    clearUsername() {
        this.username = null;
        localStorage.removeItem(this.USERNAME_KEY);
    },

    isLoggedIn() {
        return !!this.getUsername();
    },

    loadFromArray(lessons) {
        this._completed = new Set(lessons || []);
    },

    isCompleted(lessonId) {
        return this._completed.has(lessonId);
    },

    async markCompleted(lessonId) {
        this._completed.add(lessonId);
        const username = this.getUsername();
        if (username) {
            try {
                await API.markCompleted(username, lessonId);
            } catch (e) {
                console.error('Failed to save progress:', e);
            }
        }
    },

    getCompletedCount() {
        return this._completed.size;
    },

    getCompletedSet() {
        return new Set(this._completed);
    },

    async reset() {
        this._completed.clear();
        const username = this.getUsername();
        if (username) {
            try {
                await API.resetProgress(username);
            } catch (e) {
                console.error('Failed to reset progress:', e);
            }
        }
    },

    async logout() {
        this._completed.clear();
        this.clearUsername();
    },
};
