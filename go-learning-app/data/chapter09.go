package data

import "go-learning-app/models"

func loadChapter09(s *Store) {
	s.addChapter(models.Chapter{
		ID:          9,
		Title:       "テスト",
		Description: "Goの組み込みテストフレームワーク、テーブル駆動テスト、ベンチマークについて学びます。",
		Lessons: []models.LessonSummary{
			{ID: "9-1", Title: "testingパッケージ"},
			{ID: "9-2", Title: "テーブル駆動テスト"},
			{ID: "9-3", Title: "ベンチマーク"},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "9-1",
		ChapterID: 9,
		Title:     "testingパッケージ",
		Content: `Goには標準ライブラリに <code>testing</code> パッケージが含まれており、外部フレームワークなしでテストを書けます。

テストのルール:
- ファイル名は <code>_test.go</code> で終わる
- テスト関数は <code>Test</code> で始まり、<code>*testing.T</code> を引数に取る
- <code>go test</code> コマンドで実行
- <code>go test -v</code> で詳細出力
- <code>go test -cover</code> でカバレッジ表示

アサーションライブラリは標準にはないため、<code>if</code> と <code>t.Errorf()</code> を使います。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "基本的なテスト",
				Code: `// math.go
package math

func Add(a, b int) int {
    return a + b
}

func Abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}

func TestAbs(t *testing.T) {
    if Abs(-5) != 5 {
        t.Error("Abs(-5) should be 5")
    }
    if Abs(5) != 5 {
        t.Error("Abs(5) should be 5")
    }
    if Abs(0) != 0 {
        t.Error("Abs(0) should be 0")
    }
}`,
			},
			{
				Title: "テストヘルパーとサブテスト",
				Code: `package math

import "testing"

func TestAddSubtests(t *testing.T) {
    t.Run("正の数", func(t *testing.T) {
        if Add(2, 3) != 5 {
            t.Error("2+3 should be 5")
        }
    })

    t.Run("負の数", func(t *testing.T) {
        if Add(-2, -3) != -5 {
            t.Error("-2+(-3) should be -5")
        }
    })

    t.Run("ゼロ", func(t *testing.T) {
        if Add(0, 0) != 0 {
            t.Error("0+0 should be 0")
        }
    })
}`,
			},
		},
		Notes: []string{
			"_test.go ファイルはビルド時に含まれません",
			"t.Error() はテストを失敗としてマークしますが続行します",
			"t.Fatal() はテストを即座に中断します",
			"t.Run() でサブテストを作成でき、個別に実行可能です",
		},
		Exercise: &models.Exercise{
			Title:       "手動テストの実装",
			Description: "数値を2乗する関数 Square(n int) int を作成し、main関数の中で 5 の2乗が 25 になるか確認する「手動テスト」を書いてください。結果が正しければ \"PASS\"、間違っていれば \"FAIL\" と表示してください。",
			StarterCode: `package main

import "fmt"

func Square(n int) int {
    // 実装
    return 0
}

func main() {
    result := Square(5)
    expected := 25
    
    // 比較して PASS/FAIL を表示
    if result == expected {
        fmt.Println("PASS")
    } else {
        fmt.Printf("FAIL: want %d, got %d\n", expected, result)
    }
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "9-1",
		Questions: []models.Question{
			{
				ID:          "9-1-1",
				Text:        "Goのテストファイルの命名規則は？",
				Options:     []string{"test_*.go", "*_test.go", "*.test.go", "test/*.go"},
				Answer:      1,
				Explanation: "Goのテストファイルは _test.go で終わる必要があります（例: math_test.go）。",
			},
			{
				ID:          "9-1-2",
				Text:        "t.Error() と t.Fatal() の違いは？",
				Options:     []string{"同じ動作", "Error は続行、Fatal は中断", "Fatal は続行、Error は中断", "Error はログ出力のみ"},
				Answer:      1,
				Explanation: "t.Error() はテストを失敗とマークして続行しますが、t.Fatal() はテストを即座に中断します。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "9-2",
		ChapterID: 9,
		Title:     "テーブル駆動テスト",
		Content: `**テーブル駆動テスト** はGoで最も推奨されるテストパターンです。

テストケースをテーブル（スライス）として定義し、ループで実行します。これにより:
- テストケースの追加が容易
- テストコードの重複を排除
- 各ケースが独立したサブテストとして実行される

Goの標準ライブラリ自体もこのパターンを多用しています。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "テーブル駆動テスト",
				Code: `package math

import "testing"

func TestAdd_TableDriven(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"正の数同士", 2, 3, 5},
        {"負の数同士", -2, -3, -5},
        {"正と負", 5, -3, 2},
        {"ゼロ加算", 0, 5, 5},
        {"両方ゼロ", 0, 0, 0},
        {"大きな数", 1000000, 2000000, 3000000},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d",
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}`,
			},
		},
		Notes: []string{
			"テストケース構造体には name フィールドを含めるのが慣例です",
			"t.Run() の第1引数がサブテスト名になり、-run フラグで個別実行できます",
			"tt という変数名はテストケースのイディオムです（test tableの略）",
		},
		Exercise: &models.Exercise{
			Title:       "テーブル駆動テストの練習",
			Description: "数値が偶数かどうかを判定する IsEven(n int) bool 関数を作成し、複数のテストケース（スライス）を使って動作確認を行うmain関数を書いてください。",
			StarterCode: `package main

import "fmt"

func IsEven(n int) bool {
    return n % 2 == 0
}

func main() {
    tests := []struct {
        input    int
        expected bool
    }{
        {2, true},
        {3, false},
        {0, true},
        {-1, false},
    }
    
    // テーブルをループしてテスト実行
    for _, tt := range tests {
        // 結果を表示
    }
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "9-2",
		Questions: []models.Question{
			{
				ID:          "9-2-1",
				Text:        "テーブル駆動テストのメリットは？",
				Options:     []string{"実行速度が上がる", "テストケース追加が容易で重複を排除", "自動的にカバレッジ100%になる", "並列実行される"},
				Answer:      1,
				Explanation: "テーブル駆動テストは、テストケースの追加が容易でコードの重複を排除できます。",
			},
			{
				ID:          "9-2-2",
				Text:        "テストケースを個別に実行するには？",
				Options:     []string{"go test -v", "go test -run テスト名/サブテスト名", "go test -single", "go test -only テスト名"},
				Answer:      1,
				Explanation: "go test -run 'TestAdd/正の数同士' のようにして特定のサブテストだけを実行できます。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "9-3",
		ChapterID: 9,
		Title:     "ベンチマーク",
		Content: `Goの <code>testing</code> パッケージにはベンチマーク機能も含まれています。

ベンチマークのルール:
- 関数名は <code>Benchmark</code> で始まる
- <code>*testing.B</code> を引数に取る
- <code>b.N</code> 回のループを実行する（Nはランタイムが自動調整）
- <code>go test -bench=.</code> で実行

<code>b.ResetTimer()</code> でセットアップ時間をベンチマークから除外できます。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "ベンチマークの書き方",
				Code: `package math

import "testing"

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(100, 200)
    }
}

// メモリアロケーションも計測
func BenchmarkConcat(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        s := ""
        for j := 0; j < 100; j++ {
            s += "a"
        }
    }
}

// 実行: go test -bench=. -benchmem
// 出力例:
// BenchmarkAdd-8      1000000000    0.29 ns/op    0 B/op    0 allocs/op
// BenchmarkConcat-8       50000    30000 ns/op   5000 B/op   99 allocs/op`,
			},
		},
		Notes: []string{
			"b.N の値はランタイムが自動的に調整します（手動設定しない）",
			"-benchmem フラグでメモリアロケーション情報も表示されます",
			"b.ResetTimer() でセットアップ時間を除外できます",
			"b.RunParallel() で並列ベンチマークも実行できます",
		},
		Exercise: &models.Exercise{
			Title:       "ベンチマーク関数の定義",
			Description: "文字列結合を行う関数 Concat(a, b string) string を対象としたベンチマーク関数 BenchmarkConcat のコードを書いてください。（注: このエディタではベンチマークは実行できませんが、構文の練習として書いてみましょう）",
			StarterCode: `package main

import (
    "testing"
)

func Concat(a, b string) string {
    return a + b
}

// BenchmarkConcat をここに実装
func BenchmarkConcat(b *testing.B) {
    
}

func main() {
    // ベンチマークは実行できませんが、コードが正しいか確認します
    var _ func(*testing.B) = BenchmarkConcat
    fmt.Println("Code compiled successfully")
}

import "fmt"`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "9-3",
		Questions: []models.Question{
			{
				ID:          "9-3-1",
				Text:        "ベンチマーク関数の命名規則は？",
				Options:     []string{"Bench_で始まる", "Benchmark で始まる", "BM_ で始まる", "Perf で始まる"},
				Answer:      1,
				Explanation: "ベンチマーク関数はBenchmarkで始まり、*testing.B を引数に取ります。",
			},
			{
				ID:          "9-3-2",
				Text:        "b.N の値は誰が決める？",
				Options:     []string{"プログラマ", "Goランタイムが自動調整", "コンパイラ", "OS"},
				Answer:      1,
				Explanation: "b.N の値はGoランタイムが自動的に調整します。安定した計測結果が得られるまで増やされます。",
			},
		},
	})
}
