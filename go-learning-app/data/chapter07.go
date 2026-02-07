package data

import "go-learning-app/models"

func loadChapter07(s *Store) {
	s.addChapter(models.Chapter{
		ID:          7,
		Title:       "エラー処理",
		Description: "error型、カスタムエラー、panic/recoverなど、Goのエラー処理パターンを学びます。",
		Lessons: []models.LessonSummary{
			{ID: "7-1", Title: "error型"},
			{ID: "7-2", Title: "カスタムエラー"},
			{ID: "7-3", Title: "panic と recover"},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "7-1",
		ChapterID: 7,
		Title:     "error型",
		Content: `Goのエラー処理は <code>error</code> インターフェースに基づいています。

<code>error</code> は組み込みインターフェースで、<code>Error() string</code> メソッドを持ちます。

Goでは例外（try-catch）の代わりに、関数の戻り値としてエラーを返すパターンを採用しています。エラーは無視せず、必ずチェックすることが重要です。

<code>errors.New()</code> や <code>fmt.Errorf()</code> でエラーを作成できます。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "エラーハンドリングの基本",
				Code: `package main

import (
    "errors"
    "fmt"
    "strconv"
)

func parseAge(s string) (int, error) {
    age, err := strconv.Atoi(s)
    if err != nil {
        return 0, fmt.Errorf("年齢の解析に失敗: %w", err)
    }
    if age < 0 || age > 150 {
        return 0, errors.New("年齢は0から150の範囲で指定してください")
    }
    return age, nil
}

func main() {
    age, err := parseAge("25")
    if err != nil {
        fmt.Println("エラー:", err)
        return
    }
    fmt.Println("年齢:", age)

    _, err = parseAge("abc")
    if err != nil {
        fmt.Println("エラー:", err)
    }
}`,
			},
		},
		Notes: []string{
			"errorはインターフェースで、Error() string メソッドを持つ型なら何でもerrorになれます",
			"fmt.Errorf の %w 動詞でエラーをラップ（包む）できます",
			"errors.Is() と errors.As() でラップされたエラーを判定できます",
		},
		Exercise: &models.Exercise{
			Title:       "割り算のエラー",
			Description: "2つの整数を受け取り、割り算の結果を返す関数 divide を作成してください。ただし、0で割ろうとした場合はエラーを返してください。",
			StarterCode: `package main

import (
    "errors"
    "fmt"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        // エラーを返す
        
    }
    return a / b, nil
}

func main() {
    res, err := divide(10, 0)
    if err != nil {
        fmt.Println("エラー:", err)
    } else {
        fmt.Println("結果:", res)
    }
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "7-1",
		Questions: []models.Question{
			{
				ID:          "7-1-1",
				Text:        "error インターフェースが持つメソッドは？",
				Options:     []string{"String() string", "Error() string", "Message() string", "Err() string"},
				Answer:      1,
				Explanation: "error インターフェースは Error() string メソッドのみを持つシンプルなインターフェースです。",
			},
			{
				ID:          "7-1-2",
				Text:        "fmt.Errorf の %w 動詞の用途は？",
				Options:     []string{"警告を出す", "エラーをラップする", "エラーを無視する", "ログに記録する"},
				Answer:      1,
				Explanation: "%w 動詞を使うとエラーをラップ（包む）でき、errors.Is() や errors.As() でチェーンを辿れます。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "7-2",
		ChapterID: 7,
		Title:     "カスタムエラー",
		Content: `独自のエラー型を作成することで、エラーに追加情報を持たせることができます。

<code>error</code> インターフェースを実装する（<code>Error() string</code> メソッドを持つ）任意の型をエラーとして使えます。

<code>errors.Is()</code> でエラーの同一性を確認し、<code>errors.As()</code> でエラーを特定の型に変換できます。これらはラップされたエラーチェーンを辿ります。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "カスタムエラー型",
				Code: `package main

import (
    "errors"
    "fmt"
)

// カスタムエラー型
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("検証エラー [%s]: %s", e.Field, e.Message)
}

func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{
            Field:   "age",
            Message: "年齢は0以上である必要があります",
        }
    }
    return nil
}

func main() {
    err := validateAge(-5)
    if err != nil {
        // errors.As で型を判定
        var ve *ValidationError
        if errors.As(err, &ve) {
            fmt.Printf("フィールド: %s\n", ve.Field)
            fmt.Printf("メッセージ: %s\n", ve.Message)
        }
    }
}`,
			},
		},
		Notes: []string{
			"カスタムエラー型を使うとエラーに構造化データを持たせられます",
			"errors.Is() は値の同一性、errors.As() は型の一致を確認します",
			"エラー型はポインタレシーバで Error() を実装するのが一般的です",
		},
		Exercise: &models.Exercise{
			Title:       "不足エラー",
			Description: "残高(Balance)不足を表すカスタムエラー InsufficientFundsError を定義し、支払い処理(Pay)で残高不足の場合にこのエラーを返してください。",
			StarterCode: `package main

import "fmt"

// カスタムエラー定義

// 支払い関数
func Pay(balance, amount int) error {
    if balance < amount {
        // エラーを返す
        
    }
    return nil
}

func main() {
    err := Pay(100, 200)
    if err != nil {
        fmt.Println(err)
    }
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "7-2",
		Questions: []models.Question{
			{
				ID:          "7-2-1",
				Text:        "カスタムエラー型に必要なメソッドは？",
				Options:     []string{"String()", "Error() string", "Unwrap()", "Is()"},
				Answer:      1,
				Explanation: "error インターフェースを実装するには Error() string メソッドが必要です。",
			},
			{
				ID:          "7-2-2",
				Text:        "errors.As() の用途は？",
				Options:     []string{"エラーを作成する", "エラーを特定の型に変換する", "エラーをログに出す", "エラーを無視する"},
				Answer:      1,
				Explanation: "errors.As() はエラーチェーンを辿って、特定の型のエラーを取り出します。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "7-3",
		ChapterID: 7,
		Title:     "panic と recover",
		Content: `<code>panic</code> はプログラムの異常終了を引き起こす仕組みです。<code>recover</code> はpanicから回復するための仕組みです。

**panic** は本当に回復不能なエラー（プログラミングミス等）にのみ使うべきです。通常のエラーには error を使いましょう。

**recover** は <code>defer</code> 関数内でのみ機能します。panicの値をキャッチして正常な実行に戻すことができます。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "panic と recover",
				Code: `package main

import "fmt"

func safeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("回復: %v", r)
        }
    }()

    // b が 0 の場合 panic する（例示目的）
    return a / b, nil
}

func main() {
    result, err := safeDivide(10, 2)
    fmt.Printf("10/2 = %d, err = %v\n", result, err)

    result, err = safeDivide(10, 0)
    fmt.Printf("10/0 = %d, err = %v\n", result, err)
}`,
			},
			{
				Title: "deferの実行順序",
				Code: `package main

import "fmt"

func main() {
    fmt.Println("開始")

    defer fmt.Println("defer 1")
    defer fmt.Println("defer 2")
    defer fmt.Println("defer 3")

    fmt.Println("終了")
    // 出力: 開始 → 終了 → defer 3 → defer 2 → defer 1
    // (LIFO: 後入れ先出し)
}`,
			},
		},
		Notes: []string{
			"panicは通常のエラー処理には使わないでください",
			"recoverはdefer関数内でのみ機能します",
			"deferはLIFO（後入れ先出し）順で実行されます",
			"ライブラリではpanicの代わりにerrorを返すのがGoの慣例です",
		},
		Exercise: &models.Exercise{
			Title:       "パニックからの回復",
			Description: "意図的にpanicを起こす関数を作成し、それを呼び出してもメインプログラムが終了しないようにrecoverを使って回復してください。",
			StarterCode: `package main

import "fmt"

func dangerous() {
    panic("大変だ！")
}

func safeCall() {
    // deferとrecoverで回復
    
    
    dangerous()
}

func main() {
    safeCall()
    fmt.Println("メインは正常に終了")
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "7-3",
		Questions: []models.Question{
			{
				ID:          "7-3-1",
				Text:        "recoverが機能するのはどこ？",
				Options:     []string{"どこでも", "main関数内", "defer関数内のみ", "goroutine内のみ"},
				Answer:      2,
				Explanation: "recover() はdefer関数内でのみ機能します。それ以外の場所では常にnilを返します。",
			},
			{
				ID:          "7-3-2",
				Text:        "panicを使うべき場面は？",
				Options:     []string{"全てのエラー", "ファイルが見つからない時", "回復不能なプログラミングミス", "ネットワークエラー"},
				Answer:      2,
				Explanation: "panicは回復不能なエラー（初期化失敗やプログラミングミス）にのみ使うべきです。通常のエラーにはerrorを使います。",
			},
			{
				ID:          "7-3-3",
				Text:        "deferの実行順序は？",
				Options:     []string{"FIFO（先入れ先出し）", "LIFO（後入れ先出し）", "ランダム", "宣言順"},
				Answer:      1,
				Explanation: "deferはLIFO（後入れ先出し/スタック）順で実行されます。最後にdeferされたものが最初に実行されます。",
			},
		},
	})
}
