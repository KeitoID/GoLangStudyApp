package data

import "go-learning-app/models"

func loadChapter02(s *Store) {
	s.addChapter(models.Chapter{
		ID:          2,
		Title:       "制御構造",
		Description: "if文、forループ、switch文といったGoの制御構造を学びます。",
		Lessons: []models.LessonSummary{
			{ID: "2-1", Title: "if文"},
			{ID: "2-2", Title: "forループ"},
			{ID: "2-3", Title: "switch文"},
		},
	})

	// Lesson 2-1: if文
	s.addLesson(models.Lesson{
		ID:        "2-1",
		ChapterID: 2,
		Title:     "if文",
		Content: `Goの <code>if</code> 文は条件分岐を実現します。

Goのif文の特徴:
- 条件式に **括弧は不要** です（付けても動きますが非推奨）
- 波括弧 <code>{}</code> は **必須** です
- if文の条件の前に **初期化文** を書ける（スコープはifブロック内）

<code>else if</code> や <code>else</code> を使って複数の条件分岐ができます。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "基本的なif文",
				Code: `package main

import "fmt"

func main() {
    x := 10

    if x > 0 {
        fmt.Println("正の数")
    } else if x < 0 {
        fmt.Println("負の数")
    } else {
        fmt.Println("ゼロ")
    }
}`,
			},
			{
				Title: "初期化文付きif",
				Code: `package main

import (
    "fmt"
    "os"
)

func main() {
    // err のスコープは if ブロック内に限定される
    if err := doSomething(); err != nil {
        fmt.Println("エラー:", err)
        os.Exit(1)
    }
    fmt.Println("成功")
}

func doSomething() error {
    return nil
}`,
			},
		},
		Notes: []string{
			"条件式は必ず bool 型でなければなりません（0やnilは自動変換されません）",
			"初期化文付きifはエラーハンドリングでよく使われます",
		},
		Exercise: &models.Exercise{
			Title:       "偶数判定プログラム",
			Description: "変数 n に好きな整数を代入し、その数が偶数なら「偶数」、奇数なら「奇数」と表示するプログラムを書いてください。",
			StarterCode: `package main

import "fmt"

func main() {
    // 変数 n を定義
    n := 5
    
    // if文で偶数・奇数を判定して表示
    
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "2-1",
		Questions: []models.Question{
			{
				ID:          "2-1-1",
				Text:        "Goのif文で条件式を囲む括弧()は？",
				Options:     []string{"必須", "不要（非推奨）", "場合による", "エラーになる"},
				Answer:      1,
				Explanation: "Goでは条件式の括弧は不要です。付けても動作しますが、Go的なスタイルでは省略します。",
			},
			{
				ID:          "2-1-2",
				Text:        "if文の初期化文で宣言した変数のスコープは？",
				Options:     []string{"関数全体", "ifブロック内（else含む）", "ifブロックのみ（else除く）", "パッケージ全体"},
				Answer:      1,
				Explanation: "初期化文で宣言した変数はifブロックとそれに続くelse if/elseブロック内で有効です。",
			},
			{
				ID:          "2-1-3",
				Text:        "Goのif文で波括弧{}は？",
				Options:     []string{"省略可能", "一行なら省略可能", "必須", "推奨だが省略可能"},
				Answer:      2,
				Explanation: "Goではif文の波括弧は必須です。省略するとコンパイルエラーになります。",
			},
		},
	})

	// Lesson 2-2: forループ
	s.addLesson(models.Lesson{
		ID:        "2-2",
		ChapterID: 2,
		Title:     "forループ",
		Content: `Goのループは <code>for</code> のみです。while文やdo-while文はありません。

<code>for</code> は3つの形式で使えます:
1. **C言語スタイル**: <code>for i := 0; i < n; i++ { }</code>
2. **while スタイル**: <code>for 条件 { }</code>
3. **無限ループ**: <code>for { }</code>

<code>range</code> を使うとスライス、マップ、文字列などを反復処理できます。
<code>break</code> でループを抜け、<code>continue</code> で次のイテレーションに進みます。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "forの3つの形式",
				Code: `package main

import "fmt"

func main() {
    // C言語スタイル
    for i := 0; i < 5; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println()

    // whileスタイル
    n := 1
    for n < 100 {
        n *= 2
    }
    fmt.Println("n =", n)

    // 無限ループ（breakで抜ける）
    count := 0
    for {
        count++
        if count >= 3 {
            break
        }
    }
    fmt.Println("count =", count)
}`,
			},
			{
				Title: "rangeを使った反復",
				Code: `package main

import "fmt"

func main() {
    // スライスのrange
    fruits := []string{"りんご", "バナナ", "みかん"}
    for i, fruit := range fruits {
        fmt.Printf("%d: %s\n", i, fruit)
    }

    // インデックス不要な場合
    for _, fruit := range fruits {
        fmt.Println(fruit)
    }

    // 文字列のrange（runeで反復）
    for i, r := range "Go言語" {
        fmt.Printf("byte=%d, rune=%c\n", i, r)
    }
}`,
			},
		},
		Notes: []string{
			"Goにはwhileやdo-whileはなく、forだけで全てのループを表現します",
			"range は (index, value) の2つの値を返します",
			"不要な変数は _ (ブランク識別子) で捨てられます",
		},
		Exercise: &models.Exercise{
			Title:       "1から10までの合計",
			Description: "forループを使って、1から10までの整数の合計を計算して表示してください。",
			StarterCode: `package main

import "fmt"

func main() {
    sum := 0
    
    // 1から10までループして sum に加算
    
    
    fmt.Println("合計:", sum)
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "2-2",
		Questions: []models.Question{
			{
				ID:          "2-2-1",
				Text:        "Goのループ構文はいくつある？",
				Options:     []string{"for, while, doの3つ", "forのみ", "for と while の2つ", "for, while, loopの3つ"},
				Answer:      1,
				Explanation: "Goのループ構文は for のみです。whileやdo-whileの代わりにforの条件式のみ形式を使います。",
			},
			{
				ID:          "2-2-2",
				Text:        "rangeで不要な変数を無視するには？",
				Options:     []string{"null", "void", "_ (ブランク識別子)", "skip"},
				Answer:      2,
				Explanation: "_ (アンダースコア) はブランク識別子と呼ばれ、不要な値を捨てるのに使います。",
			},
			{
				ID:          "2-2-3",
				Text:        "for { } はどういうループ？",
				Options:     []string{"空のループ", "エラーになる", "無限ループ", "1回だけ実行"},
				Answer:      2,
				Explanation: "条件式を省略した for { } は無限ループになります。break で抜ける必要があります。",
			},
		},
	})

	// Lesson 2-3: switch文
	s.addLesson(models.Lesson{
		ID:        "2-3",
		ChapterID: 2,
		Title:     "switch文",
		Content: `Goの <code>switch</code> 文は強力な条件分岐を提供します。

Goのswitch文の特徴:
- **break は自動** です（fall-through しない）
- 明示的に fall-through したい場合は <code>fallthrough</code> キーワードを使う
- caseに **複数の値** を指定できる
- **式なしswitch** で複雑な条件分岐が書ける（if-else if の代替）
- <code>type switch</code> でインターフェースの型判定ができる`,
		CodeExamples: []models.CodeExample{
			{
				Title: "基本的なswitch",
				Code: `package main

import (
    "fmt"
    "time"
)

func main() {
    day := time.Now().Weekday()

    switch day {
    case time.Saturday, time.Sunday:
        fmt.Println("週末です！")
    case time.Friday:
        fmt.Println("もうすぐ週末！")
    default:
        fmt.Println("平日です")
    }
}`,
			},
			{
				Title: "式なしswitch",
				Code: `package main

import "fmt"

func main() {
    score := 85

    // 式なしswitch（if-else if の代替）
    switch {
    case score >= 90:
        fmt.Println("A")
    case score >= 80:
        fmt.Println("B")
    case score >= 70:
        fmt.Println("C")
    default:
        fmt.Println("D")
    }
}`,
			},
		},
		Notes: []string{
			"Goのswitchは自動的にbreakするため、C言語のようなfall-throughは起きません",
			"fallthrough キーワードで次のcaseに処理を落とすことができます",
			"式なしswitchは長いif-else ifチェインの代替として推奨されます",
		},
		Exercise: &models.Exercise{
			Title:       "信号機の色判定",
			Description: "変数 signal に \"red\", \"yellow\", \"blue\" のいずれかを代入し、switch文でそれぞれ「止まれ」「注意」「進め」と表示してください。それ以外は「信号機故障」と表示してください。",
			StarterCode: `package main

import "fmt"

func main() {
    signal := "red"
    
    // switch文で信号機の色に応じたメッセージを表示
    
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "2-3",
		Questions: []models.Question{
			{
				ID:          "2-3-1",
				Text:        "Goのswitchでfall-throughは？",
				Options:     []string{"自動的に起きる", "起きない（自動break）", "設定による", "常に全caseを実行"},
				Answer:      1,
				Explanation: "Goのswitchは各caseの末尾で自動的にbreakします。fall-throughさせたい場合は明示的にfallthroughと書きます。",
			},
			{
				ID:          "2-3-2",
				Text:        "式なしswitchの「switch { ... }」はどのように動作する？",
				Options:     []string{"エラーになる", "常にdefaultが実行される", "最初にtrueになるcaseが実行される", "全caseが実行される"},
				Answer:      2,
				Explanation: "式なしswitchは switch true { ... } と同等で、最初にtrueになるcaseが実行されます。",
			},
		},
	})
}
