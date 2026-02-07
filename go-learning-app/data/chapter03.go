package data

import "go-learning-app/models"

func loadChapter03(s *Store) {
	s.addChapter(models.Chapter{
		ID:          3,
		Title:       "関数とメソッド",
		Description: "関数の定義、複数戻り値、メソッド、クロージャなど、Goの関数に関する機能を学びます。",
		Lessons: []models.LessonSummary{
			{ID: "3-1", Title: "関数の定義"},
			{ID: "3-2", Title: "複数戻り値"},
			{ID: "3-3", Title: "メソッド"},
			{ID: "3-4", Title: "クロージャ"},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "3-1",
		ChapterID: 3,
		Title:     "関数の定義",
		Content: `Goの関数は <code>func</code> キーワードで定義します。

関数の基本構文: <code>func 関数名(引数) 戻り値の型 { }</code>

Goの関数は **第一級オブジェクト** です。変数に代入したり、引数として渡すことができます。

可変長引数は <code>...</code> を使って定義します。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "基本的な関数",
				Code: `package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func greet(name string) string {
    return "Hello, " + name + "!"
}

func main() {
    result := add(3, 5)
    fmt.Println(result) // 8

    msg := greet("Go")
    fmt.Println(msg) // Hello, Go!
}`,
			},
			{
				Title: "可変長引数",
				Code: `package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))       // 6
    fmt.Println(sum(1, 2, 3, 4, 5)) // 15

    // スライスを展開して渡す
    nums := []int{10, 20, 30}
    fmt.Println(sum(nums...)) // 60
}`,
			},
		},
		Notes: []string{
			"同じ型の引数は型を省略して列挙できます（例: a, b int）",
			"関数は第一級オブジェクトとして変数に代入できます",
			"可変長引数は関数内ではスライスとして扱われます",
		},
		Exercise: &models.Exercise{
			Title:       "掛け算をする関数",
			Description: "2つの整数を受け取り、その積（掛け算の結果）を返す関数 multiply を作成し、main関数から呼び出して結果を表示してください。",
			StarterCode: `package main

import "fmt"

// multiply 関数を定義
func multiply(a, b int) int {
    // ここに実装
    return 0
}

func main() {
    result := multiply(10, 20)
    fmt.Println("10 * 20 =", result)
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "3-1",
		Questions: []models.Question{
			{
				ID:          "3-1-1",
				Text:        "func add(a, b int) int の a, b の型は？",
				Options:     []string{"a はint、b は型なし", "両方ともint", "aはany、bはint", "構文エラー"},
				Answer:      1,
				Explanation: "同じ型の引数は型を省略して列挙できます。a, b int は a int, b int と同じです。",
			},
			{
				ID:          "3-1-2",
				Text:        "可変長引数 nums ...int は関数内でどう扱われる？",
				Options:     []string{"配列として", "スライスとして", "ポインタとして", "マップとして"},
				Answer:      1,
				Explanation: "可変長引数は関数内ではスライス（[]int）として扱われます。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "3-2",
		ChapterID: 3,
		Title:     "複数戻り値",
		Content: `Goの関数は **複数の値を返す** ことができます。これはGoの大きな特徴の一つです。

最も一般的なパターンは **(結果, error)** の2つの値を返すことです。これはGoのエラーハンドリングの基本パターンです。

**名前付き戻り値** を使うと、戻り値に名前を付けて関数内で変数として使えます。<code>return</code> に値を指定しない「裸のreturn」も可能ですが、短い関数以外では避けるのが推奨です。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "複数戻り値とエラーハンドリング",
				Code: `package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("ゼロ除算エラー")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 3)
    if err != nil {
        fmt.Println("エラー:", err)
        return
    }
    fmt.Printf("10 / 3 = %.2f\n", result)

    _, err = divide(10, 0)
    if err != nil {
        fmt.Println("エラー:", err)
    }
}`,
			},
			{
				Title: "名前付き戻り値",
				Code: `package main

import "fmt"

func minMax(nums []int) (min, max int) {
    min = nums[0]
    max = nums[0]
    for _, n := range nums[1:] {
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
    }
    return // 裸のreturn（min, max が返される）
}

func main() {
    lo, hi := minMax([]int{3, 1, 4, 1, 5, 9, 2, 6})
    fmt.Printf("min=%d, max=%d\n", lo, hi)
}`,
			},
		},
		Notes: []string{
			"(結果, error) パターンはGoのイディオムとして非常によく使われます",
			"不要な戻り値は _ で無視できます",
			"名前付き戻り値は関数シグネチャでドキュメントとしても機能します",
		},
		Exercise: &models.Exercise{
			Title:       "値を入れ替える関数",
			Description: "2つの文字列を受け取り、それらを入れ替えて（逆の順序で）返す関数 swap を作成してください。",
			StarterCode: `package main

import "fmt"

func swap(a, b string) (string, string) {
    // ここに実装
    return "", ""
}

func main() {
    a, b := swap("Hello", "World")
    fmt.Println(a, b) // World Hello と表示されるはず
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "3-2",
		Questions: []models.Question{
			{
				ID:          "3-2-1",
				Text:        "Goのエラーハンドリングの基本パターンは？",
				Options:     []string{"try-catch", "(結果, error) を返す", "例外をthrow", "errnoを使う"},
				Answer:      1,
				Explanation: "Goでは関数が (結果, error) の2つの値を返すパターンが標準的なエラーハンドリングです。",
			},
			{
				ID:          "3-2-2",
				Text:        "名前付き戻り値で return に値を指定しないとどうなる？",
				Options:     []string{"ゼロ値が返る", "コンパイルエラー", "名前付き変数の現在値が返る", "nil が返る"},
				Answer:      2,
				Explanation: "裸のreturnは名前付き戻り値変数の現在の値を返します。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "3-3",
		ChapterID: 3,
		Title:     "メソッド",
		Content: `メソッドは **レシーバ** を持つ関数です。特定の型に関連付けられた関数を定義できます。

レシーバには2種類あります:
- **値レシーバ**: <code>func (t Type) Method()</code> — 型のコピーを受け取る
- **ポインタレシーバ**: <code>func (t *Type) Method()</code> — 型へのポインタを受け取る

ポインタレシーバを使うと:
1. レシーバの値を変更できる
2. 大きな構造体のコピーを避けられる`,
		CodeExamples: []models.CodeExample{
			{
				Title: "メソッドの定義と使用",
				Code: `package main

import (
    "fmt"
    "math"
)

type Circle struct {
    Radius float64
}

// 値レシーバ（読み取りのみ）
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// ポインタレシーバ（値を変更可能）
func (c *Circle) Scale(factor float64) {
    c.Radius *= factor
}

func main() {
    c := Circle{Radius: 5}
    fmt.Printf("面積: %.2f\n", c.Area())

    c.Scale(2)
    fmt.Printf("スケール後の面積: %.2f\n", c.Area())
}`,
			},
		},
		Notes: []string{
			"メソッドは同じパッケージ内で定義された型にのみ追加できます",
			"レシーバの値を変更する場合はポインタレシーバを使います",
			"Goは自動的にポインタと値を相互変換してメソッドを呼び出します",
		},
		Exercise: &models.Exercise{
			Title:       "長方形の面積",
			Description: "幅(Width)と高さ(Height)を持つ構造体 Rectangle を定義し、その面積を計算して返すメソッド Area() を実装してください。",
			StarterCode: `package main

import "fmt"

// Rectangle 構造体の定義

// Area メソッドの定義

func main() {
    r := Rectangle{Width: 10, Height: 5}
    fmt.Println("面積:", r.Area())
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "3-3",
		Questions: []models.Question{
			{
				ID:          "3-3-1",
				Text:        "ポインタレシーバを使う主な理由は？",
				Options:     []string{"速度を上げるため", "レシーバの値を変更するため", "メモリを節約するため", "並行処理のため"},
				Answer:      1,
				Explanation: "ポインタレシーバの主な目的はレシーバの値を変更することです。副次的にコピーを避ける利点もあります。",
			},
			{
				ID:          "3-3-2",
				Text:        "メソッドを追加できるのは？",
				Options:     []string{"任意の型", "構造体のみ", "同じパッケージ内で定義された型", "int や string にも追加可能"},
				Answer:      2,
				Explanation: "メソッドは同じパッケージ内で定義された型にのみ追加できます。組み込み型に直接メソッドを追加することはできません。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "3-4",
		ChapterID: 3,
		Title:     "クロージャ",
		Content: `**クロージャ** は、外側のスコープの変数を参照する無名関数です。

Goでは関数は第一級オブジェクトなので、変数に代入したり、引数として渡したり、戻り値として返すことができます。

クロージャは外側の変数への **参照** を保持します（コピーではありません）。これにより状態を持つ関数を作成できます。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "クロージャの基本",
				Code: `package main

import "fmt"

func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    c1 := counter()
    fmt.Println(c1()) // 1
    fmt.Println(c1()) // 2
    fmt.Println(c1()) // 3

    // 別のカウンターは独立
    c2 := counter()
    fmt.Println(c2()) // 1
}`,
			},
			{
				Title: "関数を引数に取る",
				Code: `package main

import "fmt"

func apply(nums []int, fn func(int) int) []int {
    result := make([]int, len(nums))
    for i, n := range nums {
        result[i] = fn(n)
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5}

    doubled := apply(nums, func(n int) int {
        return n * 2
    })
    fmt.Println(doubled) // [2 4 6 8 10]

    squared := apply(nums, func(n int) int {
        return n * n
    })
    fmt.Println(squared) // [1 4 9 16 25]
}`,
			},
		},
		Notes: []string{
			"クロージャは外部変数への参照を保持します（コピーではない）",
			"ゴルーチンでクロージャを使う際はループ変数のキャプチャに注意が必要です",
			"関数型は func(引数型) 戻り値型 の形で表現します",
		},
		Exercise: &models.Exercise{
			Title:       "ステートフルなカウンター",
			Description: "呼び出すたびに指定された数だけカウントアップするクロージャを作成してください。",
			StarterCode: `package main

import "fmt"

func createAdder(step int) func() int {
    sum := 0
    return func() int {
        // ここに実装
        return sum
    }
}

func main() {
    addTwo := createAdder(2)
    fmt.Println(addTwo()) // 2
    fmt.Println(addTwo()) // 4
    fmt.Println(addTwo()) // 6
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "3-4",
		Questions: []models.Question{
			{
				ID:          "3-4-1",
				Text:        "クロージャが外部変数を保持する方法は？",
				Options:     []string{"値のコピー", "参照（ポインタ）", "グローバル変数として", "チャネル経由"},
				Answer:      1,
				Explanation: "クロージャは外部変数への参照を保持します。変数の値が変わるとクロージャ内でも反映されます。",
			},
			{
				ID:          "3-4-2",
				Text:        "func() int を返す関数の戻り値の型宣言は？",
				Options:     []string{"func int", "func() int", "function() int", "=> int"},
				Answer:      1,
				Explanation: "Goでは関数型は func(引数型) 戻り値型 の形で表現します。",
			},
		},
	})
}
