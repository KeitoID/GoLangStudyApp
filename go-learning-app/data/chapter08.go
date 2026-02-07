package data

import "go-learning-app/models"

func loadChapter08(s *Store) {
	s.addChapter(models.Chapter{
		ID:          8,
		Title:       "パッケージとモジュール",
		Description: "パッケージの構成、go modによるモジュール管理、公開/非公開のルールを学びます。",
		Lessons: []models.LessonSummary{
			{ID: "8-1", Title: "パッケージの基本"},
			{ID: "8-2", Title: "go mod"},
			{ID: "8-3", Title: "公開と非公開"},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "8-1",
		ChapterID: 8,
		Title:     "パッケージの基本",
		Content: `**パッケージ** はGoのコード管理の基本単位です。

パッケージのルール:
- 1つのディレクトリに1つのパッケージ
- 同じディレクトリ内のファイルは同じパッケージ名を宣言する
- パッケージ名はディレクトリ名と一致させるのが慣例
- <code>main</code> パッケージは実行可能プログラムのエントリポイント

<code>import</code> 文でパッケージを読み込みます。循環インポートは禁止されています。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "パッケージの構成",
				Code: `// ディレクトリ構成:
// myproject/
// ├── main.go          (package main)
// ├── math/
// │   └── math.go      (package math)
// └── utils/
//     └── helper.go    (package utils)

// --- math/math.go ---
package math

func Add(a, b int) int {
    return a + b
}

// --- main.go ---
package main

import (
    "fmt"
    "myproject/math"
)

func main() {
    result := math.Add(3, 5)
    fmt.Println(result) // 8
}`,
			},
			{
				Title: "import のバリエーション",
				Code: `package main

import (
    "fmt"                     // 標準ライブラリ
    "myproject/utils"         // 自作パッケージ

    // エイリアス
    m "myproject/math"

    // ブランクインポート（副作用のみ）
    _ "image/png"
)

func main() {
    fmt.Println(m.Add(1, 2))
    fmt.Println(utils.Helper())
}`,
			},
		},
		Notes: []string{
			"パッケージ名は短く、小文字で、一単語が推奨です",
			"循環インポート（A→B→A）はコンパイルエラーになります",
			"_ インポート（ブランクインポート）は init() 関数だけを実行するために使います",
		},
		Exercise: &models.Exercise{
			Title:       "標準パッケージの利用",
			Description: "strings パッケージと math パッケージをインポートし、\"hello world\" を大文字に変換して表示し、その後に 16 の平方根を表示するプログラムを作成してください。",
			StarterCode: `package main

import (
    "fmt"
    // strings と math をインポート
    
    
)

func main() {
    text := "hello world"
    num := 16.0
    
    // 大文字変換して表示
    
    
    // 平方根を表示
    
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "8-1",
		Questions: []models.Question{
			{
				ID:          "8-1-1",
				Text:        "1つのディレクトリに含められるパッケージの数は？",
				Options:     []string{"無制限", "1つ", "2つまで", "ファイル数と同じ"},
				Answer:      1,
				Explanation: "Goでは1つのディレクトリに1つのパッケージのみ含められます（テストファイルの _test パッケージは例外）。",
			},
			{
				ID:          "8-1-2",
				Text:        "_ \"image/png\" のブランクインポートの目的は？",
				Options:     []string{"パッケージを削除する", "init()関数のみ実行する", "テスト用", "最適化のため"},
				Answer:      1,
				Explanation: "ブランクインポートはパッケージのinit()関数を実行するためだけに使います。画像デコーダの登録などで使われます。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "8-2",
		ChapterID: 8,
		Title:     "go mod",
		Content: `**Go Modules** はGoの公式な依存関係管理システムです（Go 1.11+）。

<code>go.mod</code> ファイルでモジュールのパスと依存関係を管理します。

主要なコマンド:
- <code>go mod init モジュール名</code>: モジュールの初期化
- <code>go mod tidy</code>: 依存関係の整理（不要削除・不足追加）
- <code>go get パッケージ</code>: 依存パッケージの追加
- <code>go mod download</code>: 依存パッケージのダウンロード

<code>go.sum</code> ファイルには依存パッケージのハッシュが記録され、再現性を保証します。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "go.modの例",
				Code: `// go.mod ファイルの内容
module github.com/yourname/myproject

go 1.22

require (
    github.com/gorilla/mux v1.8.1
    golang.org/x/text v0.14.0
)`,
			},
			{
				Title: "モジュールの作成手順",
				Code: `// 1. プロジェクトディレクトリを作成
// mkdir myproject && cd myproject

// 2. モジュールを初期化
// go mod init github.com/yourname/myproject

// 3. コードを書く（main.go）
package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello!")
    })
    http.ListenAndServe(":8080", r)
}

// 4. 依存関係を解決
// go mod tidy`,
			},
		},
		Notes: []string{
			"go mod tidy は最もよく使うコマンドです",
			"go.sum は自動生成されるので手動編集は不要です",
			"モジュール名はリポジトリのURLにするのが慣例です",
		},
		Exercise: &models.Exercise{
			Title:       "モジュールの説明",
			Description: "fmtパッケージを使って、Go Modulesで依存関係を整理する際によく使うコマンド（tidy）の説明を表示するプログラムを書いてください。",
			StarterCode: `package main

import "fmt"

func main() {
    command := "go mod tidy"
    description := "依存関係を整理するコマンド"
    
    // コマンドと説明を表示
    
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "8-2",
		Questions: []models.Question{
			{
				ID:          "8-2-1",
				Text:        "go mod tidy の役割は？",
				Options:     []string{"モジュールを初期化する", "不要な依存を削除し不足を追加する", "パッケージをビルドする", "テストを実行する"},
				Answer:      1,
				Explanation: "go mod tidy は使われていない依存を削除し、不足している依存を追加します。",
			},
			{
				ID:          "8-2-2",
				Text:        "go.sum ファイルの役割は？",
				Options:     []string{"ソースコードの要約", "依存パッケージのハッシュ記録", "ビルド設定", "テスト結果"},
				Answer:      1,
				Explanation: "go.sum には依存パッケージのハッシュが記録され、ビルドの再現性を保証します。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "8-3",
		ChapterID: 8,
		Title:     "公開と非公開",
		Content: `Goでは **名前の先頭文字** でアクセス制御を行います。

- **大文字で始まる**: エクスポートされる（公開、パッケージ外からアクセス可能）
- **小文字で始まる**: エクスポートされない（非公開、パッケージ内のみ）

これは関数、変数、定数、型、構造体のフィールド、メソッドなど、全てのシンボルに適用されます。

<code>internal</code> ディレクトリを使うと、特定のモジュール内でのみアクセス可能なパッケージを作れます。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "公開と非公開の例",
				Code: `package user

import "fmt"

// User は公開型（大文字で始まる）
type User struct {
    Name  string // 公開フィールド
    Email string // 公開フィールド
    age   int    // 非公開フィールド
}

// NewUser は公開関数
func NewUser(name, email string, age int) *User {
    return &User{
        Name:  name,
        Email: email,
        age:   age,
    }
}

// GetAge は公開メソッド（非公開フィールドへのアクセサ）
func (u *User) GetAge() int {
    return u.age
}

// validate は非公開関数
func validate(email string) bool {
    return len(email) > 0
}

// String は公開メソッド
func (u *User) String() string {
    return fmt.Sprintf("%s <%s>", u.Name, u.Email)
}`,
			},
		},
		Notes: []string{
			"大文字=公開、小文字=非公開はGoの最も基本的なルールの1つです",
			"構造体のフィールドも同じルールに従います",
			"internal ディレクトリ内のパッケージは親モジュール内でのみインポート可能です",
		},
		Exercise: &models.Exercise{
			Title:       "カプセル化",
			Description: "Product構造体を定義してください。公開フィールド Name と非公開フィールド price を持ち、priceを設定する SetPrice メソッドと、priceを取得する Price メソッドを実装してください。",
			StarterCode: `package main

import "fmt"

// Product 構造体定義

// SetPrice メソッド

// Price メソッド

func main() {
    p := Product{Name: "Laptop"}
    
    // 価格を設定
    p.SetPrice(150000)
    
    // 価格を表示
    fmt.Printf("製品: %s, 価格: %d\n", p.Name, p.Price())
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "8-3",
		Questions: []models.Question{
			{
				ID:          "8-3-1",
				Text:        "Goでシンボルを公開するには？",
				Options:     []string{"public キーワードを付ける", "名前を大文字で始める", "export する", "アノテーションを付ける"},
				Answer:      1,
				Explanation: "Goでは名前を大文字で始めるだけでエクスポート（公開）されます。キーワードは不要です。",
			},
			{
				ID:          "8-3-2",
				Text:        "internal ディレクトリ内のパッケージの特徴は？",
				Options:     []string{"テスト専用", "親モジュール内でのみインポート可能", "自動的に公開される", "ビルドされない"},
				Answer:      1,
				Explanation: "internal ディレクトリ内のパッケージは、そのinternalディレクトリの親以下からのみインポートできます。",
			},
		},
	})
}
