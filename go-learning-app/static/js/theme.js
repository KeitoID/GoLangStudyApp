// Dark/Light theme toggle
const Theme = {
    STORAGE_KEY: 'go-learning-theme',

    init() {
        const saved = localStorage.getItem(this.STORAGE_KEY);
        const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
        const theme = saved || (prefersDark ? 'dark' : 'light');
        this.apply(theme);

        document.getElementById('themeToggle').addEventListener('click', () => {
            const current = document.documentElement.getAttribute('data-theme');
            this.apply(current === 'dark' ? 'light' : 'dark');
        });
    },

    apply(theme) {
        document.documentElement.setAttribute('data-theme', theme);
        localStorage.setItem(this.STORAGE_KEY, theme);
        document.getElementById('themeIcon').textContent = theme === 'dark' ? '\u2600\uFE0F' : '\uD83C\uDF19';
    }
};
