package data

import "go-learning-app/models"

func loadChapter01(s *Store) {
	s.addChapter(models.Chapter{
		ID:          1,
		Title:       "Go言語の基礎",
		Description: "Go言語の基本的な構文と概念を学びます。Hello Worldから始めて、変数、データ型、fmtパッケージの使い方を習得します。",
		Lessons: []models.LessonSummary{
			{ID: "1-1", Title: "Hello World"},
			{ID: "1-2", Title: "変数と定数"},
			{ID: "1-3", Title: "基本データ型"},
			{ID: "1-4", Title: "fmtパッケージ"},
		},
	})

	// Lesson 1-1: Hello World
	s.addLesson(models.Lesson{
		ID:        "1-1",
		ChapterID: 1,
		Title:     "Hello World",
		Content: `Go言語での最初のプログラムを書いてみましょう。

Goのプログラムは必ず **package宣言** から始まります。実行可能なプログラムは <code>package main</code> を使い、エントリポイントとして <code>main()</code> 関数を定義します。

<code>import</code> 文で標準ライブラリやパッケージを読み込みます。<code>fmt</code> パッケージは書式付き出力を提供する、最もよく使うパッケージの一つです。

Goのプログラムを実行するには <code>go run ファイル名.go</code> コマンドを使います。コンパイルして実行ファイルを作るには <code>go build</code> を使います。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "最初のGoプログラム",
				Code: `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`,
			},
			{
				Title: "複数のimport",
				Code: `package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("現在時刻:", time.Now())
}`,
			},
		},
		Notes: []string{
			"Goではセミコロンは不要です（コンパイラが自動挿入します）",
			"未使用のimportはコンパイルエラーになります",
			"main パッケージの main() 関数がプログラムのエントリポイントです",
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "1-1",
		Questions: []models.Question{
			{
				ID:      "1-1-1",
				Text:    "Goの実行可能プログラムで必要なパッケージ名は？",
				Options: []string{"main", "app", "go", "run"},
				Answer:  0,
				Explanation: "実行可能なGoプログラムは必ず package main を宣言する必要があります。",
			},
			{
				ID:      "1-1-2",
				Text:    "Goで未使用のimportがあるとどうなる？",
				Options: []string{"警告が出る", "無視される", "コンパイルエラーになる", "実行時エラーになる"},
				Answer:  2,
				Explanation: "Goでは未使用のimportはコンパイルエラーになります。これはコードの清潔さを保つための設計方針です。",
			},
			{
				ID:      "1-1-3",
				Text:    "fmt.Println() の役割は？",
				Options: []string{"ファイルに書き込む", "標準出力に改行付きで表示する", "エラーを出力する", "ログに記録する"},
				Answer:  1,
				Explanation: "fmt.Println() は引数を標準出力に表示し、最後に改行を追加します。",
			},
		},
	})

	// Lesson 1-2: 変数と定数
	s.addLesson(models.Lesson{
		ID:        "1-2",
		ChapterID: 1,
		Title:     "変数と定数",
		Content: `Goにおける変数宣言と定数の使い方を学びます。

Goでは変数を宣言する方法が複数あります:
- <code>var</code> キーワードを使った宣言
- <code>:=</code> (短縮変数宣言) を使った宣言（関数内のみ）

変数は宣言時に **ゼロ値** で初期化されます。数値型は <code>0</code>、文字列型は <code>""</code>、bool型は <code>false</code> です。

定数は <code>const</code> キーワードで宣言し、コンパイル時に値が決定される必要があります。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "変数宣言の方法",
				Code: `package main

import "fmt"

func main() {
    // var を使った宣言
    var name string = "Go"
    var age int = 15

    // 型推論（型を省略）
    var language = "Go言語"

    // 短縮変数宣言（関数内のみ）
    message := "Hello!"

    // 複数変数の同時宣言
    var x, y int = 10, 20

    fmt.Println(name, age, language, message, x, y)
}`,
			},
			{
				Title: "ゼロ値と定数",
				Code: `package main

import "fmt"

const Pi = 3.14159
const (
    StatusOK    = 200
    StatusError = 500
)

func main() {
    // ゼロ値の確認
    var i int       // 0
    var f float64   // 0.0
    var b bool      // false
    var s string    // ""

    fmt.Printf("int: %d, float: %f, bool: %t, string: %q\n", i, f, b, s)
    fmt.Println("Pi =", Pi)
}`,
			},
		},
		Notes: []string{
			":= は関数の外では使えません",
			"Goでは未使用の変数もコンパイルエラーになります",
			"定数には := は使えません（const を使います）",
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "1-2",
		Questions: []models.Question{
			{
				ID:      "1-2-1",
				Text:    ":= （短縮変数宣言）はどこで使える？",
				Options: []string{"どこでも使える", "関数内のみ", "パッケージレベルのみ", "main関数内のみ"},
				Answer:  1,
				Explanation: ":= は関数の内部でのみ使用可能です。パッケージレベルでは var を使う必要があります。",
			},
			{
				ID:      "1-2-2",
				Text:    "int型のゼロ値は？",
				Options: []string{"nil", "0", "false", "\"\""},
				Answer:  1,
				Explanation: "int型のゼロ値は 0 です。各型にはそれぞれのゼロ値があります。",
			},
			{
				ID:      "1-2-3",
				Text:    "Goで未使用の変数があるとどうなる？",
				Options: []string{"警告が出る", "自動的に削除される", "コンパイルエラーになる", "何も起きない"},
				Answer:  2,
				Explanation: "Goでは未使用の変数はコンパイルエラーになります。コードの清潔さを保つための設計です。",
			},
		},
	})

	// Lesson 1-3: 基本データ型
	s.addLesson(models.Lesson{
		ID:        "1-3",
		ChapterID: 1,
		Title:     "基本データ型",
		Content: `Goの基本的なデータ型について学びます。

Goは **静的型付け言語** です。主要なデータ型は以下の通りです:

**整数型**: <code>int</code>, <code>int8</code>, <code>int16</code>, <code>int32</code>, <code>int64</code> (符号なしは <code>uint</code> 系)
**浮動小数点型**: <code>float32</code>, <code>float64</code>
**論理型**: <code>bool</code>
**文字列型**: <code>string</code> (UTF-8、イミュータブル)
**バイト型**: <code>byte</code> (<code>uint8</code> のエイリアス)
**ルーン型**: <code>rune</code> (<code>int32</code> のエイリアス、Unicode コードポイント)

型変換は **明示的** に行う必要があります。暗黙の型変換は行われません。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "基本データ型の使用",
				Code: `package main

import "fmt"

func main() {
    // 整数型
    var age int = 30
    var small int8 = 127  // -128 ~ 127

    // 浮動小数点型
    var pi float64 = 3.14159
    var e float32 = 2.718

    // 論理型
    var isGo bool = true

    // 文字列型
    var greeting string = "こんにちは"

    // rune（Unicode文字）
    var r rune = '漢'

    fmt.Printf("age: %d, small: %d\n", age, small)
    fmt.Printf("pi: %f, e: %f\n", pi, e)
    fmt.Printf("isGo: %t\n", isGo)
    fmt.Printf("greeting: %s (len=%d bytes)\n", greeting, len(greeting))
    fmt.Printf("rune: %c (Unicode: U+%04X)\n", r, r)
}`,
			},
			{
				Title: "型変換",
				Code: `package main

import "fmt"

func main() {
    // 明示的な型変換が必要
    var i int = 42
    var f float64 = float64(i)
    var u uint = uint(f)

    fmt.Println(i, f, u)

    // 文字列とバイト列の変換
    s := "Hello, Go!"
    b := []byte(s)
    s2 := string(b)
    fmt.Println(s, b, s2)
}`,
			},
		},
		Notes: []string{
			"int のサイズはプラットフォームにより32ビットまたは64ビットです",
			"string は UTF-8 エンコードされたバイト列です",
			"暗黙の型変換はないため、異なる型同士の演算にはキャストが必要です",
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "1-3",
		Questions: []models.Question{
			{
				ID:      "1-3-1",
				Text:    "runeは何のエイリアス？",
				Options: []string{"uint8", "int32", "int64", "byte"},
				Answer:  1,
				Explanation: "rune は int32 のエイリアスで、Unicodeコードポイントを表します。",
			},
			{
				ID:      "1-3-2",
				Text:    "Goで int を float64 に変換するには？",
				Options: []string{"自動変換される", "float64(i)", "(float64)i", "i.toFloat64()"},
				Answer:  1,
				Explanation: "Goでは明示的な型変換が必要です。float64(i) の形で変換します。",
			},
		},
	})

	// Lesson 1-4: fmtパッケージ
	s.addLesson(models.Lesson{
		ID:        "1-4",
		ChapterID: 1,
		Title:     "fmtパッケージ",
		Content: `<code>fmt</code> パッケージはGoの書式付きI/Oを提供する標準パッケージです。

主要な出力関数:
- <code>fmt.Print()</code>: 改行なしで出力
- <code>fmt.Println()</code>: 改行付きで出力
- <code>fmt.Printf()</code>: 書式指定付きで出力

主要な書式指定子:
- <code>%d</code>: 整数
- <code>%f</code>: 浮動小数点数
- <code>%s</code>: 文字列
- <code>%t</code>: 真偽値
- <code>%v</code>: デフォルト形式
- <code>%T</code>: 型名
- <code>%q</code>: クォート付き文字列

入力関数: <code>fmt.Scan()</code>, <code>fmt.Scanf()</code>, <code>fmt.Scanln()</code>
文字列生成: <code>fmt.Sprintf()</code> (文字列を返す), <code>fmt.Fprintf()</code> (io.Writerに書き込む)`,
		CodeExamples: []models.CodeExample{
			{
				Title: "出力関数の使い分け",
				Code: `package main

import "fmt"

func main() {
    name := "Go"
    version := 1.22

    // Print: 改行なし
    fmt.Print("Hello, ")
    fmt.Print(name)
    fmt.Println() // 改行だけ

    // Println: スペース区切り＋改行
    fmt.Println("言語:", name, "バージョン:", version)

    // Printf: 書式指定
    fmt.Printf("言語: %s, バージョン: %.2f\n", name, version)
}`,
			},
			{
				Title: "Sprintfと各種フォーマット",
				Code: `package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {
    p := Point{10, 20}

    // Sprintf: 文字列として返す
    s := fmt.Sprintf("座標: (%d, %d)", p.X, p.Y)
    fmt.Println(s)

    // %v: デフォルト形式
    fmt.Printf("%%v:  %v\n", p)
    // %+v: フィールド名付き
    fmt.Printf("%%+v: %+v\n", p)
    // %#v: Go構文形式
    fmt.Printf("%%#v: %#v\n", p)
    // %T: 型名
    fmt.Printf("%%T:  %T\n", p)
}`,
			},
		},
		Notes: []string{
			"fmt.Sprintf() は文字列を返すだけで出力しません",
			"%%  でリテラルの % を出力できます",
			"fmt.Errorf() でエラー値を書式付きで生成できます",
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "1-4",
		Questions: []models.Question{
			{
				ID:      "1-4-1",
				Text:    "fmt.Sprintf() は何を返す？",
				Options: []string{"int", "error", "string", "[]byte"},
				Answer:  2,
				Explanation: "fmt.Sprintf() は書式付きの文字列を返します。標準出力には出力しません。",
			},
			{
				ID:      "1-4-2",
				Text:    "%v フォーマット指定子の意味は？",
				Options: []string{"verbose出力", "デフォルト形式で値を表示", "バージョン表示", "検証モード"},
				Answer:  1,
				Explanation: "%v はデフォルト形式で値を表示します。構造体の場合はフィールド値が表示されます。",
			},
			{
				ID:      "1-4-3",
				Text:    "変数の型名を表示する書式指定子は？",
				Options: []string{"%v", "%s", "%T", "%t"},
				Answer:  2,
				Explanation: "%T は値の型名を表示します。デバッグ時に型を確認するのに便利です。",
			},
		},
	})
}
