package data

import "go-learning-app/models"

func loadChapter04(s *Store) {
	s.addChapter(models.Chapter{
		ID:          4,
		Title:       "データ構造",
		Description: "配列、スライス、マップといったGoの基本的なデータ構造を学びます。",
		Lessons: []models.LessonSummary{
			{ID: "4-1", Title: "配列"},
			{ID: "4-2", Title: "スライス"},
			{ID: "4-3", Title: "マップ"},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "4-1",
		ChapterID: 4,
		Title:     "配列",
		Content: `**配列** は固定長の同じ型の要素の集合です。

Goの配列の特徴:
- サイズは型の一部です（<code>[3]int</code> と <code>[5]int</code> は異なる型）
- 値型です（代入や関数引数ではコピーされます）
- サイズはコンパイル時に決定される必要があります

実際のGoプログラムでは配列よりもスライスの方がよく使われます。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "配列の基本",
				Code: `package main

import "fmt"

func main() {
    // 配列の宣言
    var a [3]int
    a[0] = 10
    a[1] = 20
    a[2] = 30
    fmt.Println(a) // [10 20 30]

    // リテラルで初期化
    b := [3]string{"Go", "Python", "Rust"}
    fmt.Println(b)

    // サイズを自動推論
    c := [...]int{1, 2, 3, 4, 5}
    fmt.Println(len(c)) // 5

    // 配列は値型（コピーされる）
    d := a
    d[0] = 999
    fmt.Println(a) // [10 20 30] （変更されない）
    fmt.Println(d) // [999 20 30]
}`,
			},
		},
		Notes: []string{
			"配列のサイズは型の一部なので、[3]int と [5]int は別の型です",
			"[...]型{値} でコンパイラにサイズを推論させることができます",
			"実務ではスライスの方がはるかに多く使われます",
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "4-1",
		Questions: []models.Question{
			{
				ID:      "4-1-1",
				Text:    "[3]int と [5]int は同じ型か？",
				Options: []string{"同じ型", "異なる型", "場合による", "互換性がある"},
				Answer:  1,
				Explanation: "配列のサイズは型の一部なので、[3]int と [5]int は異なる型です。",
			},
			{
				ID:      "4-1-2",
				Text:    "配列を別の変数に代入するとどうなる？",
				Options: []string{"参照が共有される", "全要素がコピーされる", "ポインタがコピーされる", "エラーになる"},
				Answer:  1,
				Explanation: "Goの配列は値型なので、代入すると全要素がコピーされます。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "4-2",
		ChapterID: 4,
		Title:     "スライス",
		Content: `**スライス** は可変長の配列への参照です。Goで最もよく使われるデータ構造です。

スライスは3つの要素で構成されます:
- **ポインタ**: 基底配列の要素を指す
- **長さ (len)**: スライスの要素数
- **容量 (cap)**: 基底配列のスライス開始位置からの要素数

<code>make()</code> で作成し、<code>append()</code> で要素を追加します。容量が足りなくなると自動的に拡張されます。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "スライスの基本操作",
				Code: `package main

import "fmt"

func main() {
    // スライスリテラル
    s := []int{1, 2, 3, 4, 5}
    fmt.Println(s) // [1 2 3 4 5]

    // make で作成
    s2 := make([]int, 3, 10) // len=3, cap=10
    fmt.Printf("len=%d, cap=%d\n", len(s2), cap(s2))

    // append で要素追加
    s2 = append(s2, 4, 5, 6)
    fmt.Println(s2) // [0 0 0 4 5 6]

    // スライス式
    sub := s[1:4] // インデックス1から3まで
    fmt.Println(sub) // [2 3 4]
}`,
			},
			{
				Title: "スライスの注意点",
				Code: `package main

import "fmt"

func main() {
    // スライスは参照型
    original := []int{1, 2, 3}
    copied := original
    copied[0] = 999
    fmt.Println(original) // [999 2 3] 変更が反映される！

    // 独立したコピーを作る
    independent := make([]int, len(original))
    copy(independent, original)
    independent[0] = 1
    fmt.Println(original)    // [999 2 3]
    fmt.Println(independent) // [1 2 3]
}`,
			},
		},
		Notes: []string{
			"スライスは参照型のため、代入すると基底配列を共有します",
			"独立したコピーが必要な場合は copy() を使います",
			"append は容量超過時に新しい基底配列を確保します",
			"nil スライスと空スライスは異なりますが、len()はどちらも0です",
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "4-2",
		Questions: []models.Question{
			{
				ID:      "4-2-1",
				Text:    "スライスを構成する3つの要素は？",
				Options: []string{"型、値、サイズ", "ポインタ、長さ、容量", "インデックス、値、長さ", "配列、開始、終了"},
				Answer:  1,
				Explanation: "スライスはポインタ（基底配列への参照）、長さ（len）、容量（cap）で構成されます。",
			},
			{
				ID:      "4-2-2",
				Text:    "s[1:4] はどの要素を含む？",
				Options: []string{"インデックス1,2,3,4", "インデックス1,2,3", "インデックス0,1,2,3", "インデックス1,2"},
				Answer:  1,
				Explanation: "スライス式 s[1:4] はインデックス1から3まで（4は含まない）の要素を含みます。",
			},
			{
				ID:      "4-2-3",
				Text:    "スライスの独立したコピーを作るには？",
				Options: []string{"= で代入", "copy() を使う", "clone() を使う", "new() を使う"},
				Answer:  1,
				Explanation: "copy() 関数を使うと、スライスの要素を別のスライスにコピーできます。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "4-3",
		ChapterID: 4,
		Title:     "マップ",
		Content: `**マップ** はキーと値のペアを格納するデータ構造です（他の言語の辞書やハッシュマップに相当）。

マップの特徴:
- <code>make(map[キー型]値型)</code> で作成
- キーは比較可能な型（==で比較できる型）である必要がある
- 存在しないキーを読むとゼロ値が返る
- 2つ目の戻り値で存在確認ができる（comma ok イディオム）
- <code>delete()</code> で要素を削除`,
		CodeExamples: []models.CodeExample{
			{
				Title: "マップの基本操作",
				Code: `package main

import "fmt"

func main() {
    // マップの作成
    ages := map[string]int{
        "Alice": 30,
        "Bob":   25,
    }

    // 要素の追加・更新
    ages["Charlie"] = 35

    // 要素の取得
    fmt.Println("Alice:", ages["Alice"])

    // 存在確認（comma ok イディオム）
    age, ok := ages["Dave"]
    if ok {
        fmt.Println("Dave:", age)
    } else {
        fmt.Println("Dave は存在しません")
    }

    // 要素の削除
    delete(ages, "Bob")

    // 反復処理
    for name, age := range ages {
        fmt.Printf("%s: %d歳\n", name, age)
    }

    fmt.Println("人数:", len(ages))
}`,
			},
		},
		Notes: []string{
			"マップの反復順序は保証されません（毎回異なる可能性があります）",
			"nil マップへの書き込みは panic を起こします（必ず make で初期化）",
			"マップは並行処理で安全ではありません（sync.Map を検討してください）",
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "4-3",
		Questions: []models.Question{
			{
				ID:      "4-3-1",
				Text:    "マップに存在しないキーを読むとどうなる？",
				Options: []string{"panic が起きる", "nil が返る", "ゼロ値が返る", "エラーが返る"},
				Answer:  2,
				Explanation: "存在しないキーを読むとその型のゼロ値が返ります。存在確認には comma ok イディオムを使います。",
			},
			{
				ID:      "4-3-2",
				Text:    "マップの反復順序は？",
				Options: []string{"挿入順", "キーの昇順", "保証されない", "キーの降順"},
				Answer:  2,
				Explanation: "マップの反復順序は保証されていません。順序が必要な場合はキーをソートする必要があります。",
			},
			{
				ID:      "4-3-3",
				Text:    "age, ok := m[\"key\"] の ok は何を表す？",
				Options: []string{"値が正しいか", "キーが存在するか", "型が一致するか", "マップが初期化されているか"},
				Answer:  1,
				Explanation: "comma ok イディオムの2番目の戻り値は、キーが存在するかどうかを bool で返します。",
			},
		},
	})
}
