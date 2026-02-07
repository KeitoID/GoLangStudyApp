// UI rendering components
const Components = {
    chaptersData: [],
    totalLessons: 0,

    renderSidebar(chapters) {
        this.chaptersData = chapters;
        this.totalLessons = chapters.reduce((sum, ch) => sum + ch.lessons.length, 0);
        const nav = document.getElementById('sidebarNav');
        const completed = Progress.getCompletedSet();

        let html = '';
        for (const ch of chapters) {
            const isOpen = this._isChapterOpen(ch.id);
            html += `<div class="chapter-group" data-chapter="${ch.id}">
                <div class="chapter-header" onclick="App.toggleChapter(${ch.id})">
                    <span class="chapter-num">${ch.id}</span>
                    <span class="chapter-title">${ch.title}</span>
                    <span class="chapter-toggle ${isOpen ? 'open' : ''}">\u25B6</span>
                </div>
                <div class="lesson-list ${isOpen ? 'open' : ''}">`;

            for (const lesson of ch.lessons) {
                const isComplete = completed.has(lesson.id);
                const isActive = App.currentLessonId === lesson.id;
                html += `<div class="lesson-item ${isActive ? 'active' : ''}"
                              onclick="App.navigateTo('${lesson.id}')">
                    <span class="lesson-check ${isComplete ? 'completed' : ''}"></span>
                    <span>${lesson.title}</span>
                </div>`;
            }

            html += `</div></div>`;
        }

        nav.innerHTML = html;
        this.updateProgress();
    },

    _isChapterOpen(chapterId) {
        const key = `go-learning-ch-${chapterId}`;
        return sessionStorage.getItem(key) !== 'closed';
    },

    toggleChapter(chapterId) {
        const key = `go-learning-ch-${chapterId}`;
        const group = document.querySelector(`[data-chapter="${chapterId}"]`);
        const list = group.querySelector('.lesson-list');
        const toggle = group.querySelector('.chapter-toggle');
        const isOpen = list.classList.contains('open');

        if (isOpen) {
            list.classList.remove('open');
            toggle.classList.remove('open');
            sessionStorage.setItem(key, 'closed');
        } else {
            list.classList.add('open');
            toggle.classList.add('open');
            sessionStorage.removeItem(key);
        }
    },

    updateProgress() {
        const completedCount = Progress.getCompletedCount();
        const total = this.totalLessons;
        const pct = total > 0 ? Math.round((completedCount / total) * 100) : 0;

        document.getElementById('progressText').textContent = `${completedCount}/${total}`;
        document.getElementById('progressFill').style.width = `${pct}%`;
    },

    updateSidebarActive(lessonId) {
        document.querySelectorAll('.lesson-item').forEach(el => {
            el.classList.remove('active');
        });
        document.querySelectorAll('.lesson-item').forEach(el => {
            if (el.getAttribute('onclick')?.includes(`'${lessonId}'`)) {
                el.classList.add('active');
            }
        });
    },

    renderLesson(lesson) {
        const view = document.getElementById('lessonView');
        const chapter = this.chaptersData.find(ch => ch.id === lesson.chapterId);
        const chapterTitle = chapter ? chapter.title : '';

        // Parse content - convert line breaks to paragraphs
        const contentHtml = lesson.content
            .split('\n\n')
            .map(p => `<p>${p.trim()}</p>`)
            .join('');

        let html = `
            <div class="lesson-breadcrumb">第${lesson.chapterId}章: ${chapterTitle}</div>
            <h1 class="lesson-title">${lesson.title}</h1>
            <div class="lesson-content">${contentHtml}</div>`;

        // Code examples with "Try it" button
        for (let i = 0; i < lesson.codeExamples.length; i++) {
            const ex = lesson.codeExamples[i];
            html += `
            <div class="code-example">
                <div class="code-example-title">${ex.title}</div>
                <pre class="language-go"><code class="language-go">${this._escapeHtml(ex.code)}</code></pre>
                <div class="code-example-actions">
                    <button class="try-btn" onclick="Components.loadCodeToEditor('${lesson.id}', ${i})">
                        ▶ 試してみる
                    </button>
                </div>
            </div>`;
        }

        // Notes
        if (lesson.notes && lesson.notes.length > 0) {
            html += `
            <div class="lesson-notes">
                <h3>ポイント</h3>
                <ul>
                    ${lesson.notes.map(n => `<li>${n}</li>`).join('')}
                </ul>
            </div>`;
        }

        // Exercise section with editor
        const starterCode = lesson.exercise?.starterCode || lesson.codeExamples[0]?.code || `package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}`;

        html += `
        <div class="exercise-section">
            <div class="exercise-header">
                <span class="exercise-title">💻 ${lesson.exercise?.title || 'コードを書いてみよう'}</span>
            </div>
            <p class="exercise-description">${lesson.exercise?.description || '上記のコードを参考に、自分でコードを書いて実行してみましょう。'}</p>
            <div id="editorMount"></div>
        </div>`;

        // Quiz button
        html += `
            <button class="quiz-start-btn" onclick="App.startQuiz('${lesson.id}')">
                クイズに挑戦する
            </button>`;

        // Navigation
        html += this._renderLessonNav(lesson.id);

        view.innerHTML = html;

        // Initialize editor
        const editorMount = document.getElementById('editorMount');
        if (editorMount && typeof Editor !== 'undefined') {
            Editor.init(editorMount, starterCode);
        }

        // Store lesson for later use
        this._currentLesson = lesson;

        // Highlight code
        if (window.Prism) {
            Prism.highlightAllUnder(view);
        }
    },

    // Load code example into editor
    loadCodeToEditor(lessonId, exampleIndex) {
        if (this._currentLesson && this._currentLesson.id === lessonId) {
            const code = this._currentLesson.codeExamples[exampleIndex]?.code;
            if (code && Editor.editor) {
                Editor.currentStarterCode = code;
                Editor.editor.setValue(code);
                // Scroll to editor
                document.getElementById('editorMount')?.scrollIntoView({ behavior: 'smooth' });
            }
        }
    },

    _renderLessonNav(currentId) {
        const allLessons = [];
        for (const ch of this.chaptersData) {
            for (const l of ch.lessons) {
                allLessons.push(l);
            }
        }

        const idx = allLessons.findIndex(l => l.id === currentId);
        const prev = idx > 0 ? allLessons[idx - 1] : null;
        const next = idx < allLessons.length - 1 ? allLessons[idx + 1] : null;

        let html = '<div class="lesson-nav">';

        if (prev) {
            html += `<button class="lesson-nav-btn" onclick="App.navigateTo('${prev.id}')">
                \u2190 ${prev.title}
            </button>`;
        } else {
            html += '<div></div>';
        }

        if (next) {
            html += `<button class="lesson-nav-btn" onclick="App.navigateTo('${next.id}')">
                ${next.title} \u2192
            </button>`;
        } else {
            html += '<div></div>';
        }

        html += '</div>';
        return html;
    },

    renderQuiz(quiz) {
        const view = document.getElementById('quizView');
        const questions = quiz.questions;

        let html = `
            <div class="quiz-header">
                <h2>クイズ</h2>
                <span class="quiz-progress-text">${questions.length}問</span>
            </div>`;

        for (let i = 0; i < questions.length; i++) {
            const q = questions[i];
            const selected = Quiz.getSelectedAnswer(q.id);

            html += `
            <div class="quiz-question" data-question="${q.id}">
                <div class="quiz-question-text">Q${i + 1}. ${q.text}</div>
                <div class="quiz-options">`;

            const labels = ['A', 'B', 'C', 'D'];
            for (let j = 0; j < q.options.length; j++) {
                let classes = 'quiz-option';
                if (Quiz.submitted) {
                    classes += ' disabled';
                    if (j === q.answer) classes += ' correct';
                    else if (j === selected && j !== q.answer) classes += ' wrong';
                } else if (j === selected) {
                    classes += ' selected';
                }

                html += `
                    <div class="${classes}" onclick="App.selectQuizOption('${q.id}', ${j})">
                        <span class="quiz-option-marker">${labels[j]}</span>
                        <span>${q.options[j]}</span>
                    </div>`;
            }

            html += `</div>
                <div class="quiz-explanation ${Quiz.submitted ? 'show' : ''}" id="explanation-${q.id}">
                    ${Quiz.submitted ? (Quiz.isCorrect(q.id) ? '\u2705 正解! ' : '\u274C 不正解. ') + q.explanation : q.explanation}
                </div>
            </div>`;
        }

        // Actions
        if (!Quiz.submitted) {
            html += `
            <div class="quiz-actions">
                <button class="btn btn-primary" onclick="App.submitQuiz()" ${Quiz.allAnswered() ? '' : 'disabled'}>
                    回答を送信
                </button>
                <button class="btn btn-secondary" onclick="App.showLesson()">
                    レッスンに戻る
                </button>
            </div>`;
        }

        view.innerHTML = html;
    },

    renderQuizResult(result, lessonId) {
        const view = document.getElementById('quizView');

        // Keep existing questions, add result at top
        const resultHtml = `
            <div class="quiz-result">
                <div class="quiz-result-score ${result.passed ? 'pass' : 'fail'}">
                    ${result.percent}%
                </div>
                <div class="quiz-result-message">
                    ${result.passed ? 'おめでとうございます！合格です！' : '惜しい！もう一度挑戦してみましょう。'}
                </div>
                <div class="quiz-result-detail">
                    ${result.total}問中${result.correct}問正解 (合格ライン: 70%)
                </div>
                <div class="quiz-actions" style="justify-content:center;">
                    ${result.passed
                ? `<button class="btn btn-primary" onclick="App.showLesson()">レッスンに戻る</button>`
                : `<button class="btn btn-primary" onclick="App.startQuiz('${lessonId}')">もう一度挑戦</button>
                           <button class="btn btn-secondary" onclick="App.showLesson()">レッスンに戻る</button>`
            }
                </div>
            </div>`;

        // Re-render quiz with results, prepending the score
        this.renderQuiz(Quiz.currentQuiz);
        view.insertAdjacentHTML('afterbegin', resultHtml);
    },

    _escapeHtml(str) {
        const div = document.createElement('div');
        div.textContent = str;
        return div.innerHTML;
    },

    showView(viewName) {
        document.getElementById('welcomeScreen').style.display = viewName === 'welcome' ? '' : 'none';
        document.getElementById('lessonView').style.display = viewName === 'lesson' ? '' : 'none';
        document.getElementById('quizView').style.display = viewName === 'quiz' ? '' : 'none';
    }
};
