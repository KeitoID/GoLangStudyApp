// Progress management using localStorage
const Progress = {
    STORAGE_KEY: 'go-learning-progress',

    _load() {
        try {
            const data = localStorage.getItem(this.STORAGE_KEY);
            return data ? JSON.parse(data) : {};
        } catch {
            return {};
        }
    },

    _save(data) {
        localStorage.setItem(this.STORAGE_KEY, JSON.stringify(data));
    },

    isCompleted(lessonId) {
        const data = this._load();
        return data[lessonId] === true;
    },

    markCompleted(lessonId) {
        const data = this._load();
        data[lessonId] = true;
        this._save(data);
    },

    getCompletedCount() {
        const data = this._load();
        return Object.values(data).filter(v => v === true).length;
    },

    getCompletedSet() {
        const data = this._load();
        const set = new Set();
        for (const [k, v] of Object.entries(data)) {
            if (v === true) set.add(k);
        }
        return set;
    },

    reset() {
        localStorage.removeItem(this.STORAGE_KEY);
    }
};
