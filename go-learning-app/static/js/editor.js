// Code Editor component using CodeMirror
const Editor = {
    editor: null,
    currentStarterCode: '',
    isRunning: false,

    // Initialize CodeMirror editor
    init(container, starterCode = '') {
        this.currentStarterCode = starterCode;
        
        const editorContainer = document.createElement('div');
        editorContainer.className = 'editor-container';
        
        editorContainer.innerHTML = `
            <div class="editor-header">
                <span class="editor-title">📝 コードエディタ</span>
                <div class="editor-actions">
                    <button class="editor-btn reset-btn" onclick="Editor.reset()" title="リセット">
                        🔄 リセット
                    </button>
                    <button class="editor-btn run-btn" onclick="Editor.run()" title="実行 (Ctrl+Enter)">
                        ▶ 実行
                    </button>
                </div>
            </div>
            <div class="editor-wrapper">
                <textarea id="codeEditor">${this._escapeHtml(starterCode)}</textarea>
            </div>
            <div class="output-container">
                <div class="output-header">
                    <span>📤 出力</span>
                    <span class="output-status" id="outputStatus"></span>
                </div>
                <pre class="output-content" id="outputContent">実行ボタンを押してコードを実行してください</pre>
            </div>
        `;
        
        container.appendChild(editorContainer);
        
        // Initialize CodeMirror
        const textarea = document.getElementById('codeEditor');
        this.editor = CodeMirror.fromTextArea(textarea, {
            mode: 'go',
            theme: document.documentElement.getAttribute('data-theme') === 'dark' ? 'material-darker' : 'default',
            lineNumbers: true,
            indentUnit: 4,
            tabSize: 4,
            indentWithTabs: true,
            lineWrapping: true,
            autoCloseBrackets: true,
            matchBrackets: true,
            extraKeys: {
                'Ctrl-Enter': () => this.run(),
                'Cmd-Enter': () => this.run(),
                'Tab': (cm) => {
                    if (cm.somethingSelected()) {
                        cm.indentSelection('add');
                    } else {
                        cm.replaceSelection('\t', 'end');
                    }
                }
            }
        });
        
        // Set initial size
        this.editor.setSize('100%', '300px');
    },

    // Run the code
    async run() {
        if (this.isRunning || !this.editor) return;
        
        const code = this.editor.getValue();
        if (!code.trim()) {
            this._showOutput('コードを入力してください', true);
            return;
        }
        
        this.isRunning = true;
        this._setRunning(true);
        
        try {
            const response = await fetch('/api/run', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ code })
            });
            
            const result = await response.json();
            
            if (result.error) {
                this._showOutput(result.output || result.error, true);
            } else {
                this._showOutput(result.output || '(出力なし)', false);
            }
        } catch (err) {
            this._showOutput('実行エラー: ' + err.message, true);
        } finally {
            this.isRunning = false;
            this._setRunning(false);
        }
    },

    // Reset to starter code
    reset() {
        if (this.editor) {
            this.editor.setValue(this.currentStarterCode);
            document.getElementById('outputContent').textContent = '実行ボタンを押してコードを実行してください';
            document.getElementById('outputContent').classList.remove('error');
            document.getElementById('outputStatus').textContent = '';
        }
    },

    // Update theme when toggled
    updateTheme(isDark) {
        if (this.editor) {
            this.editor.setOption('theme', isDark ? 'material-darker' : 'default');
        }
    },

    // Private methods
    _setRunning(running) {
        const btn = document.querySelector('.run-btn');
        const status = document.getElementById('outputStatus');
        
        if (running) {
            btn.disabled = true;
            btn.innerHTML = '⏳ 実行中...';
            status.textContent = '実行中...';
        } else {
            btn.disabled = false;
            btn.innerHTML = '▶ 実行';
        }
    },

    _showOutput(text, isError = false) {
        const output = document.getElementById('outputContent');
        const status = document.getElementById('outputStatus');
        
        output.textContent = text;
        output.classList.toggle('error', isError);
        status.textContent = isError ? '❌ エラー' : '✅ 完了';
    },

    _escapeHtml(str) {
        const div = document.createElement('div');
        div.textContent = str;
        return div.innerHTML;
    }
};
