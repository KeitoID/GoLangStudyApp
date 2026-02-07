package data

import "go-learning-app/models"

func loadChapter05(s *Store) {
	s.addChapter(models.Chapter{
		ID:          5,
		Title:       "構造体とインターフェース",
		Description: "構造体の定義、インターフェースの概念、型アサーション、構造体の埋め込みを学びます。",
		Lessons: []models.LessonSummary{
			{ID: "5-1", Title: "構造体"},
			{ID: "5-2", Title: "インターフェース"},
			{ID: "5-3", Title: "型アサーション"},
			{ID: "5-4", Title: "構造体の埋め込み"},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "5-1",
		ChapterID: 5,
		Title:     "構造体",
		Content: `**構造体 (struct)** はフィールドの集合で、データをグループ化するための型です。

Goにはクラスがありませんが、構造体とメソッドを組み合わせてオブジェクト指向的な設計ができます。

構造体は <code>type 名前 struct { }</code> で定義します。フィールドの先頭が大文字なら公開（エクスポート）、小文字なら非公開です。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "構造体の定義と使用",
				Code: `package main

import "fmt"

type User struct {
    Name  string
    Email string
    Age   int
}

func main() {
    // 構造体の初期化
    u1 := User{Name: "Alice", Email: "alice@example.com", Age: 30}

    // フィールドアクセス
    fmt.Println(u1.Name)

    // ポインタ
    u2 := &User{Name: "Bob", Age: 25}
    u2.Email = "bob@example.com" // ポインタでも . でアクセス可能
    fmt.Printf("%+v\n", u2)

    // ゼロ値で初期化
    var u3 User
    fmt.Printf("%+v\n", u3) // {Name: Email: Age:0}
}`,
			},
			{
				Title: "コンストラクタパターン",
				Code: `package main

import "fmt"

type Server struct {
    Host string
    Port int
}

// コンストラクタ関数（Goの慣例: New + 型名）
func NewServer(host string, port int) *Server {
    return &Server{
        Host: host,
        Port: port,
    }
}

func (s *Server) Address() string {
    return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func main() {
    srv := NewServer("localhost", 8080)
    fmt.Println(srv.Address())
}`,
			},
		},
		Notes: []string{
			"Goにはクラスやコンストラクタはなく、構造体 + New関数パターンを使います",
			"フィールド名が大文字で始まるとパッケージ外からアクセス可能です",
			"ポインタ経由でもドット記法でフィールドにアクセスできます",
		},
		Exercise: &models.Exercise{
			Title:       "書籍データの構造体",
			Description: "タイトル(Title)、著者(Author)、価格(Price)を持つ構造体 Book を定義し、好きな本のデータを作成して内容を表示してください。",
			StarterCode: `package main

import "fmt"

// Book 構造体を定義

func main() {
    // Bookのインスタンスを作成
    
    // 内容を表示
    fmt.Printf("タイトル: %s, 著者: %s, 価格: %d円\n", )
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "5-1",
		Questions: []models.Question{
			{
				ID:          "5-1-1",
				Text:        "構造体のフィールドを外部パッケージからアクセス可能にするには？",
				Options:     []string{"public キーワードを付ける", "フィールド名を大文字で始める", "export キーワードを付ける", "アクセス修飾子を設定する"},
				Answer:      1,
				Explanation: "Goでは名前が大文字で始まるとエクスポート（公開）されます。これは全てのシンボルに共通のルールです。",
			},
			{
				ID:          "5-1-2",
				Text:        "Goのコンストラクタの慣例は？",
				Options:     []string{"init() メソッド", "constructor() 関数", "New + 型名の関数", "__init__ メソッド"},
				Answer:      2,
				Explanation: "Goでは New + 型名（例: NewServer）のパターンがコンストラクタの慣例です。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "5-2",
		ChapterID: 5,
		Title:     "インターフェース",
		Content: `**インターフェース** はメソッドシグネチャの集合を定義する型です。

Goのインターフェースの最大の特徴は **暗黙的実装** です。型がインターフェースの全メソッドを実装していれば、明示的な宣言なしにそのインターフェースを満たします。

<code>interface{}</code>（空インターフェース）は任意の型の値を保持できます。Go 1.18 以降は <code>any</code> というエイリアスが使えます。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "インターフェースの定義と実装",
				Code: `package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

func printShape(s Shape) {
    fmt.Printf("面積: %.2f, 周囲: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    r := Rectangle{Width: 10, Height: 5}
    c := Circle{Radius: 7}

    printShape(r) // Rectangle は Shape を暗黙的に実装
    printShape(c) // Circle も Shape を暗黙的に実装
}`,
			},
		},
		Notes: []string{
			"Goのインターフェースは暗黙的に実装されます（implements キーワードは不要）",
			"小さなインターフェースが推奨されます（io.Reader, io.Writer など）",
			"any は interface{} のエイリアスです（Go 1.18+）",
		},
		Exercise: &models.Exercise{
			Title:       "動物の鳴き声",
			Description: "Speak() string メソッドを持つインターフェース Animal を定義し、Dog（犬）と Cat（猫）の構造体にそれぞれ実装して、鳴き声（\"ワンワン\", \"ニャー\"）を表示させてください。",
			StarterCode: `package main

import "fmt"

// Animal インターフェースの定義

// Dog 構造体と Speak メソッド

// Cat 構造体と Speak メソッド

func main() {
    var animals []Animal
    // DogとCatを追加
    
    // ループでSpeakを呼び出す
    for _, a := range animals {
        fmt.Println(a.Speak())
    }
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "5-2",
		Questions: []models.Question{
			{
				ID:          "5-2-1",
				Text:        "Goでインターフェースを実装するには？",
				Options:     []string{"implements キーワード", "全メソッドを実装するだけ（暗黙的）", "register() を呼ぶ", "@ アノテーション"},
				Answer:      1,
				Explanation: "Goではインターフェースの全メソッドを実装するだけで、自動的にそのインターフェースを満たします。",
			},
			{
				ID:          "5-2-2",
				Text:        "空インターフェース interface{} の特徴は？",
				Options:     []string{"何も保持できない", "任意の型の値を保持できる", "構造体のみ保持できる", "nil のみ保持できる"},
				Answer:      1,
				Explanation: "空インターフェースはメソッドが0個なので、全ての型が暗黙的に実装しています。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "5-3",
		ChapterID: 5,
		Title:     "型アサーション",
		Content: `**型アサーション** は、インターフェース値の具体的な型にアクセスする方法です。

構文: <code>value := i.(Type)</code>

型アサーションが失敗すると panic が起きます。安全に行うには **comma ok パターン** を使います: <code>value, ok := i.(Type)</code>

**型スイッチ** を使うと、複数の型を判定できます。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "型アサーションと型スイッチ",
				Code: `package main

import "fmt"

func describe(i any) string {
    // 型スイッチ
    switch v := i.(type) {
    case int:
        return fmt.Sprintf("整数: %d", v)
    case string:
        return fmt.Sprintf("文字列: %q (長さ%d)", v, len(v))
    case bool:
        return fmt.Sprintf("真偽値: %t", v)
    case []int:
        return fmt.Sprintf("intスライス: %v", v)
    default:
        return fmt.Sprintf("不明な型: %T", v)
    }
}

func main() {
    fmt.Println(describe(42))
    fmt.Println(describe("hello"))
    fmt.Println(describe(true))
    fmt.Println(describe([]int{1, 2, 3}))

    // comma ok パターン
    var i any = "Go言語"
    s, ok := i.(string)
    if ok {
        fmt.Println("文字列:", s)
    }
}`,
			},
		},
		Notes: []string{
			"型アサーションはインターフェース型の値に対してのみ使用できます",
			"comma ok パターンを使わない型アサーションは失敗時に panic します",
			"型スイッチの case では変数 v に具体的な型の値が代入されます",
		},
		Exercise: &models.Exercise{
			Title:       "型の判別",
			Description: "any型の引数を受け取り、それが int なら2倍の値を、string なら \"Hello, \" + 文字列 を表示し、それ以外なら \"Unknown type\" と表示する関数 process を作成してください。",
			StarterCode: `package main

import "fmt"

func process(v any) {
    // 型スイッチで処理を分岐
    
}

func main() {
    process(10)
    process("World")
    process(true)
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "5-3",
		Questions: []models.Question{
			{
				ID:          "5-3-1",
				Text:        "型アサーション i.(string) が失敗するとどうなる？",
				Options:     []string{"nil が返る", "panic が起きる", "空文字列が返る", "コンパイルエラー"},
				Answer:      1,
				Explanation: "comma ok パターンを使わない型アサーションは、失敗時に panic を起こします。",
			},
			{
				ID:          "5-3-2",
				Text:        "型スイッチで使うキーワードは？",
				Options:     []string{"i.(type)", "typeof(i)", "i.type()", "reflect.TypeOf(i)"},
				Answer:      0,
				Explanation: "型スイッチでは switch v := i.(type) の形式を使います。(type) は switch 文内でのみ使えます。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "5-4",
		ChapterID: 5,
		Title:     "構造体の埋め込み",
		Content: `**構造体の埋め込み (Embedding)** は、Goで構成（コンポジション）を実現する方法です。

フィールド名を省略して型だけを指定すると、その型のフィールドとメソッドが昇格（プロモート）されます。これにより継承に似た効果が得られますが、あくまでコンポジション（組み合わせ）です。

インターフェースも埋め込むことができ、大きなインターフェースを小さなインターフェースの組み合わせで定義できます。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "構造体の埋め込み",
				Code: `package main

import "fmt"

type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return a.Name + "が鳴いています"
}

type Dog struct {
    Animal // 埋め込み（フィールド名なし）
    Breed  string
}

func main() {
    d := Dog{
        Animal: Animal{Name: "ポチ"},
        Breed:  "柴犬",
    }

    // Animal のフィールドに直接アクセス
    fmt.Println(d.Name)    // ポチ
    fmt.Println(d.Speak()) // ポチが鳴いています
    fmt.Println(d.Breed)   // 柴犬
}`,
			},
			{
				Title: "インターフェースの埋め込み",
				Code: `package main

import "fmt"

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Reader と Writer を埋め込んで組み合わせ
type ReadWriter interface {
    Reader
    Writer
}

type MyReadWriter struct{}

func (rw MyReadWriter) Read(p []byte) (int, error) {
    fmt.Println("Reading...")
    return 0, nil
}

func (rw MyReadWriter) Write(p []byte) (int, error) {
    fmt.Println("Writing...")
    return len(p), nil
}

func main() {
    var rw ReadWriter = MyReadWriter{}
    rw.Read(nil)
    rw.Write([]byte("hello"))
}`,
			},
		},
		Notes: []string{
			"埋め込みは継承ではなくコンポジションです",
			"埋め込まれた型のメソッドは昇格して直接呼び出せます",
			"io.ReadWriter は io.Reader と io.Writer の埋め込みで定義されています",
		},
		Exercise: &models.Exercise{
			Title:       "プログラマーの定義",
			Description: "Nameを持つ Person 構造体を定義し、それを埋め込んだ Programmer 構造体（Languageフィールドを追加）を作成してください。Programmerのインスタンスを作成し、NameとLanguageを表示してください。",
			StarterCode: `package main

import "fmt"

// Person 構造体

// Programmer 構造体（Personを埋め込み）

func main() {
    // Programmerのインスタンス作成
    
    // フィールドを表示（Nameは昇格しているため直接アクセス可）
    
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "5-4",
		Questions: []models.Question{
			{
				ID:          "5-4-1",
				Text:        "構造体の埋め込みは何を実現する？",
				Options:     []string{"継承", "コンポジション（組み合わせ）", "ポリモーフィズム", "カプセル化"},
				Answer:      1,
				Explanation: "Goの構造体の埋め込みはコンポジション（組み合わせ）を実現します。継承ではありません。",
			},
			{
				ID:          "5-4-2",
				Text:        "埋め込まれた型のメソッドはどうなる？",
				Options:     []string{"呼べなくなる", "直接呼び出せる（昇格）", "オーバーライドされる", "別名で呼ぶ必要がある"},
				Answer:      1,
				Explanation: "埋め込まれた型のメソッドは昇格（プロモート）され、外側の型から直接呼び出せます。",
			},
		},
	})
}
