// Main application: initialization, routing, state management
const App = {
    currentLessonId: null,
    currentLesson: null,

    async init() {
        Theme.init();
        this._setupMobileMenu();

        // Check login state
        if (Progress.isLoggedIn()) {
            await this._loginAndLoad(Progress.getUsername());
        } else {
            this._showLoginModal();
            return;
        }

        // Hash-based routing
        window.addEventListener('hashchange', () => this._handleRoute());
        this._handleRoute();
    },

    async _loginAndLoad(username) {
        try {
            const result = await API.login(username);
            Progress.setUsername(result.username);
            Progress.loadFromArray(result.progress);
            this._updateUsernameDisplay(result.username);
        } catch (e) {
            console.error('Login failed:', e);
            Progress.clearUsername();
            this._showLoginModal();
            return;
        }

        try {
            const chapters = await API.getChapters();
            Components.renderSidebar(chapters);
        } catch (e) {
            console.error('Failed to load chapters:', e);
        }

        this._hideLoginModal();
    },

    _showLoginModal() {
        const modal = document.getElementById('loginModal');
        if (modal) modal.style.display = 'flex';
    },

    _hideLoginModal() {
        const modal = document.getElementById('loginModal');
        if (modal) modal.style.display = 'none';
    },

    async handleLogin() {
        const input = document.getElementById('usernameInput');
        const errorEl = document.getElementById('loginError');
        const username = input.value.trim();

        if (!username) {
            errorEl.textContent = '\u30e6\u30fc\u30b6\u30fc\u540d\u3092\u5165\u529b\u3057\u3066\u304f\u3060\u3055\u3044';
            return;
        }

        errorEl.textContent = '';

        try {
            await this._loginAndLoad(username);

            // Set up routing after successful login
            window.addEventListener('hashchange', () => this._handleRoute());
            this._handleRoute();
        } catch (e) {
            errorEl.textContent = '\u30ed\u30b0\u30a4\u30f3\u306b\u5931\u6557\u3057\u307e\u3057\u305f';
        }
    },

    handleLogout() {
        Progress.logout();
        this._updateUsernameDisplay(null);
        // Reload to show login modal
        window.location.hash = '';
        window.location.reload();
    },

    _updateUsernameDisplay(username) {
        const el = document.getElementById('usernameDisplay');
        if (el) {
            el.textContent = username || '';
            el.style.display = username ? '' : 'none';
        }
    },

    _handleRoute() {
        const hash = window.location.hash.slice(1); // remove #
        if (hash.startsWith('lesson/')) {
            const id = hash.replace('lesson/', '');
            this.navigateTo(id, false);
        } else if (hash === '' || hash === '/') {
            Components.showView('welcome');
            this.currentLessonId = null;
        }
    },

    async navigateTo(lessonId, updateHash = true) {
        if (updateHash) {
            window.location.hash = `lesson/${lessonId}`;
            // hashchange handler will call navigateTo again with false
            return;
        }

        try {
            this.currentLessonId = lessonId;
            this.currentLesson = await API.getLesson(lessonId);
            Components.renderLesson(this.currentLesson);
            Components.showView('lesson');
            Components.updateSidebarActive(lessonId);
            // Close mobile sidebar
            this._closeMobileSidebar();
            // Scroll to top
            window.scrollTo(0, 0);
        } catch (e) {
            console.error('Failed to load lesson:', e);
        }
    },

    async startQuiz(lessonId) {
        const quiz = await Quiz.load(lessonId);
        if (!quiz) return;
        Components.renderQuiz(quiz);
        Components.showView('quiz');
        window.scrollTo(0, 0);
    },

    selectQuizOption(questionId, optionIndex) {
        if (Quiz.submitted) return;
        Quiz.selectAnswer(questionId, optionIndex);
        Components.renderQuiz(Quiz.currentQuiz);
    },

    submitQuiz() {
        const result = Quiz.submit();
        if (!result) return;

        Components.renderQuizResult(result, this.currentLessonId);
        Components.updateProgress();
        // Re-render sidebar to update checkmarks
        if (Components.chaptersData.length > 0) {
            Components.renderSidebar(Components.chaptersData);
            Components.updateSidebarActive(this.currentLessonId);
        }
        window.scrollTo(0, 0);
    },

    showLesson() {
        if (this.currentLesson) {
            Components.renderLesson(this.currentLesson);
            Components.showView('lesson');
        }
    },

    toggleChapter(chapterId) {
        Components.toggleChapter(chapterId);
    },

    _setupMobileMenu() {
        const toggle = document.getElementById('menuToggle');
        const sidebar = document.getElementById('sidebar');
        const overlay = document.getElementById('sidebarOverlay');

        toggle.addEventListener('click', () => {
            sidebar.classList.toggle('open');
            overlay.classList.toggle('open');
        });

        overlay.addEventListener('click', () => {
            this._closeMobileSidebar();
        });
    },

    _closeMobileSidebar() {
        document.getElementById('sidebar').classList.remove('open');
        document.getElementById('sidebarOverlay').classList.remove('open');
    }
};

// Start the application
document.addEventListener('DOMContentLoaded', () => App.init());
