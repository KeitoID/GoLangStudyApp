// Quiz logic and scoring
const Quiz = {
    currentQuiz: null,
    answers: {},
    submitted: false,

    async load(lessonId) {
        this.answers = {};
        this.submitted = false;
        try {
            this.currentQuiz = await API.getQuiz(lessonId);
            return this.currentQuiz;
        } catch {
            this.currentQuiz = null;
            return null;
        }
    },

    selectAnswer(questionId, optionIndex) {
        if (this.submitted) return;
        this.answers[questionId] = optionIndex;
    },

    getSelectedAnswer(questionId) {
        return this.answers[questionId] !== undefined ? this.answers[questionId] : -1;
    },

    allAnswered() {
        if (!this.currentQuiz) return false;
        return this.currentQuiz.questions.every(q => this.answers[q.id] !== undefined);
    },

    submit() {
        if (!this.currentQuiz || !this.allAnswered()) return null;
        this.submitted = true;

        const questions = this.currentQuiz.questions;
        let correct = 0;

        for (const q of questions) {
            if (this.answers[q.id] === q.answer) {
                correct++;
            }
        }

        const total = questions.length;
        const percent = Math.round((correct / total) * 100);
        const passed = percent >= 70;

        if (passed) {
            Progress.markCompleted(this.currentQuiz.lessonId);
        }

        return { correct, total, percent, passed };
    },

    isCorrect(questionId) {
        if (!this.currentQuiz) return false;
        const q = this.currentQuiz.questions.find(q => q.id === questionId);
        return q && this.answers[questionId] === q.answer;
    },

    getCorrectAnswer(questionId) {
        if (!this.currentQuiz) return -1;
        const q = this.currentQuiz.questions.find(q => q.id === questionId);
        return q ? q.answer : -1;
    }
};
