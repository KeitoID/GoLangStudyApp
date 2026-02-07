package data

import "go-learning-app/models"

func loadChapter10(s *Store) {
	s.addChapter(models.Chapter{
		ID:          10,
		Title:       "実践パターン",
		Description: "CLIツール、HTTPサーバー、JSON/APIの実装パターンを通して実践的なGoプログラミングを学びます。",
		Lessons: []models.LessonSummary{
			{ID: "10-1", Title: "CLIツール"},
			{ID: "10-2", Title: "HTTPサーバー"},
			{ID: "10-3", Title: "JSON と API"},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "10-1",
		ChapterID: 10,
		Title:     "CLIツール",
		Content: `GoはCLI (Command Line Interface) ツールの開発に非常に適しています。シングルバイナリにコンパイルされるため、配布が容易です。

標準ライブラリで使えるCLI関連パッケージ:
- <code>os</code>: コマンドライン引数 (<code>os.Args</code>)、環境変数、プロセス制御
- <code>flag</code>: コマンドラインフラグのパース
- <code>bufio</code>: バッファ付き入力（標準入力の読み取り）
- <code>os/exec</code>: 外部コマンドの実行`,
		CodeExamples: []models.CodeExample{
			{
				Title: "flagパッケージの使用",
				Code: `package main

import (
    "flag"
    "fmt"
)

func main() {
    // フラグの定義
    name := flag.String("name", "World", "挨拶する相手")
    count := flag.Int("count", 1, "繰り返し回数")
    upper := flag.Bool("upper", false, "大文字に変換")

    // フラグのパース
    flag.Parse()

    // フラグの使用
    for i := 0; i < *count; i++ {
        msg := fmt.Sprintf("Hello, %s!", *name)
        if *upper {
            msg = fmt.Sprintf("HELLO, %s!", *name)
        }
        fmt.Println(msg)
    }

    // 残りの引数
    fmt.Println("残りの引数:", flag.Args())
}

// 実行例:
// go run main.go -name=Go -count=3
// go run main.go -upper -name=World`,
			},
			{
				Title: "標準入力の読み取り",
				Code: `package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("名前を入力: ")
    if scanner.Scan() {
        name := strings.TrimSpace(scanner.Text())
        fmt.Printf("こんにちは、%sさん！\n", name)
    }

    // 複数行の入力
    fmt.Println("テキストを入力（Ctrl+Dで終了）:")
    lines := 0
    for scanner.Scan() {
        lines++
    }
    fmt.Printf("%d行読み込みました\n", lines)
}`,
			},
		},
		Notes: []string{
			"flag パッケージはポインタを返すので、使用時は * で参照します",
			"Go製のCLIツールはシングルバイナリなのでDocker等での配布が容易です",
			"大規模なCLIには cobra や urfave/cli などのサードパーティライブラリも人気です",
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "10-1",
		Questions: []models.Question{
			{
				ID:      "10-1-1",
				Text:    "flag.String() の戻り値の型は？",
				Options: []string{"string", "*string", "[]string", "flag.Value"},
				Answer:  1,
				Explanation: "flag.String() は *string（文字列へのポインタ）を返します。使用時は * で値を取り出します。",
			},
			{
				ID:      "10-1-2",
				Text:    "GoのCLIツールの配布が容易な理由は？",
				Options: []string{"インタプリタ言語だから", "シングルバイナリにコンパイルされるから", "JVMで動くから", "Dockerが必須だから"},
				Answer:  1,
				Explanation: "Goは依存ライブラリも含めてシングルバイナリにコンパイルされるため、実行環境に依存しません。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "10-2",
		ChapterID: 10,
		Title:     "HTTPサーバー",
		Content: `Goの標準ライブラリ <code>net/http</code> パッケージだけで本格的なHTTPサーバーを構築できます。

主要コンポーネント:
- <code>http.HandleFunc()</code>: ハンドラー関数の登録
- <code>http.ListenAndServe()</code>: サーバーの起動
- <code>http.ServeMux</code>: ルーター（Go 1.22でパスパラメータ対応）
- <code>http.Handler</code> インターフェース: カスタムハンドラー

Go 1.22 以降、<code>http.ServeMux</code> がメソッドベースルーティングとパスパラメータをサポートしました。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "基本的なHTTPサーバー",
				Code: `package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    // 基本的なハンドラー
    mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!")
    })

    // パスパラメータ（Go 1.22+）
    mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
        id := r.PathValue("id")
        fmt.Fprintf(w, "User ID: %s\n", id)
    })

    // ミドルウェア
    handler := loggingMiddleware(mux)

    fmt.Println("サーバー起動: http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}`,
			},
		},
		Notes: []string{
			"Go 1.22 から ServeMux でメソッドとパスパラメータが使えます",
			"net/http だけで本番レベルのサーバーが構築できます",
			"ミドルウェアは http.Handler をラップするパターンで実装します",
			"http.Server 構造体でタイムアウトなどの詳細設定が可能です",
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "10-2",
		Questions: []models.Question{
			{
				ID:      "10-2-1",
				Text:    "Go 1.22 の ServeMux の新機能は？",
				Options: []string{"WebSocket対応", "メソッドベースルーティングとパスパラメータ", "自動HTTPS", "GraphQL対応"},
				Answer:  1,
				Explanation: "Go 1.22からServeMuxがメソッドベースルーティング（GET /pathなど）とパスパラメータ（{id}）をサポートします。",
			},
			{
				ID:      "10-2-2",
				Text:    "r.PathValue(\"id\") は何を返す？",
				Options: []string{"クエリパラメータ", "URLパスパラメータの値", "ヘッダーの値", "フォームの値"},
				Answer:  1,
				Explanation: "r.PathValue() はURLパスに定義したパラメータ（例: /users/{id}）の値を返します。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "10-3",
		ChapterID: 10,
		Title:     "JSON と API",
		Content: `<code>encoding/json</code> パッケージでJSONのエンコード/デコードができます。

主要関数:
- <code>json.Marshal()</code>: Go → JSON バイト列
- <code>json.Unmarshal()</code>: JSON バイト列 → Go
- <code>json.NewEncoder()</code>: io.Writer に直接エンコード
- <code>json.NewDecoder()</code>: io.Reader から直接デコード

構造体タグ (<code>json:"フィールド名"</code>) でJSONフィールド名をカスタマイズできます。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "JSONのエンコードとデコード",
				Code: `package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type User struct {
    ID       int    ` + "`json:\"id\"`" + `
    Name     string ` + "`json:\"name\"`" + `
    Email    string ` + "`json:\"email\"`" + `
    Password string ` + "`json:\"-\"`" + `           // JSON出力から除外
    Age      int    ` + "`json:\"age,omitempty\"`" + ` // ゼロ値なら省略
}

func main() {
    // エンコード（Go → JSON）
    user := User{
        ID:       1,
        Name:     "Alice",
        Email:    "alice@example.com",
        Password: "secret",
    }

    data, err := json.MarshalIndent(user, "", "  ")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(data))
    // {"id":1,"name":"Alice","email":"alice@example.com"}
    // Password は除外、Age は omitempty で省略

    // デコード（JSON → Go）
    jsonStr := ` + "`" + `{"id":2,"name":"Bob","email":"bob@example.com","age":25}` + "`" + `
    var user2 User
    err = json.Unmarshal([]byte(jsonStr), &user2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", user2)
}`,
			},
			{
				Title: "JSON APIエンドポイント",
				Code: `package main

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Status  string ` + "`json:\"status\"`" + `
    Message string ` + "`json:\"message\"`" + `
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    // リクエストボディのデコード
    var req struct {
        Name string ` + "`json:\"name\"`" + `
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{
            Status: "error", Message: "不正なJSON",
        })
        return
    }

    // レスポンスの返却
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Response{
        Status:  "ok",
        Message: "Hello, " + req.Name,
    })
}`,
			},
		},
		Notes: []string{
			"json:\"-\" でフィールドをJSONから除外できます（パスワード等に使用）",
			"omitempty でゼロ値のフィールドをJSON出力から省略できます",
			"json.NewEncoder/Decoder は io.Writer/Reader と直接やり取りします",
			"構造体タグはバッククォートで囲みます",
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "10-3",
		Questions: []models.Question{
			{
				ID:      "10-3-1",
				Text:    "json:\"-\" タグの意味は？",
				Options: []string{"フィールドを必須にする", "フィールドをJSONから除外する", "フィールド名をハイフンにする", "デフォルト値を設定する"},
				Answer:  1,
				Explanation: "json:\"-\" はそのフィールドをJSONのエンコード/デコードから完全に除外します。",
			},
			{
				ID:      "10-3-2",
				Text:    "omitempty の効果は？",
				Options: []string{"必須フィールドにする", "常に出力する", "ゼロ値なら出力を省略する", "null を出力する"},
				Answer:  2,
				Explanation: "omitempty はフィールドがゼロ値（0, \"\", nil等）の場合、JSON出力から省略します。",
			},
			{
				ID:      "10-3-3",
				Text:    "json.NewEncoder(w).Encode(v) の利点は？",
				Options: []string{"高速になる", "io.Writerに直接書き込める", "エラーが出ない", "自動でgzip圧縮される"},
				Answer:  1,
				Explanation: "NewEncoderはio.Writer（http.ResponseWriter等）に直接JSONを書き込めるため、中間のバイト列を作る必要がありません。",
			},
		},
	})
}
