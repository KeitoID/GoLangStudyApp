package data

import "go-learning-app/models"

func loadChapter06(s *Store) {
	s.addChapter(models.Chapter{
		ID:          6,
		Title:       "並行処理",
		Description: "ゴルーチン、チャネル、select文、syncパッケージなど、Goの並行処理機能を学びます。",
		Lessons: []models.LessonSummary{
			{ID: "6-1", Title: "ゴルーチン"},
			{ID: "6-2", Title: "チャネル"},
			{ID: "6-3", Title: "select文"},
			{ID: "6-4", Title: "syncパッケージ"},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "6-1",
		ChapterID: 6,
		Title:     "ゴルーチン",
		Content: `**ゴルーチン (goroutine)** はGoの軽量スレッドです。<code>go</code> キーワードを付けて関数を呼び出すだけで並行実行できます。

ゴルーチンの特徴:
- OSスレッドよりもはるかに軽量（初期スタックサイズ約2KB）
- Goランタイムが管理するグリーンスレッド
- 数千〜数百万のゴルーチンを同時に実行可能
- <code>go func()</code> で起動する

メインのゴルーチンが終了すると、全てのゴルーチンも終了します。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "ゴルーチンの基本",
				Code: `package main

import (
    "fmt"
    "time"
)

func sayHello(name string) {
    for i := 0; i < 3; i++ {
        fmt.Printf("Hello, %s! (%d)\n", name, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // ゴルーチンとして起動
    go sayHello("Go")
    go sayHello("World")

    // 無名関数のゴルーチン
    go func() {
        fmt.Println("無名関数のゴルーチン")
    }()

    // メインが終了するとゴルーチンも終了する
    time.Sleep(500 * time.Millisecond)
    fmt.Println("メイン終了")
}`,
			},
		},
		Notes: []string{
			"ゴルーチンは go キーワードで簡単に起動できます",
			"メインゴルーチンが終了すると全ゴルーチンが強制終了します",
			"実際のプログラムでは time.Sleep ではなく sync.WaitGroup やチャネルで同期します",
		},
		Exercise: &models.Exercise{
			Title:       "並行カウントダウン",
			Description: "ゴルーチンを使って、3から1までのカウントダウンを1秒間隔で表示する関数を起動してください。メイン関数では3.5秒待機して終了してください。",
			StarterCode: `package main

import (
    "fmt"
    "time"
)

func countdown() {
    // 3から1までループ
    for i := 3; i > 0; i-- {
        fmt.Println(i)
        // 1秒待機
        time.Sleep(time.Second)
    }
}

func main() {
    // ゴルーチン起動
    
    // メインゴルーチンで待機
    fmt.Println("Start")
    time.Sleep(3500 * time.Millisecond)
    fmt.Println("Finish")
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "6-1",
		Questions: []models.Question{
			{
				ID:          "6-1-1",
				Text:        "ゴルーチンを起動するキーワードは？",
				Options:     []string{"async", "thread", "go", "spawn"},
				Answer:      2,
				Explanation: "go キーワードを関数呼び出しの前に付けるだけでゴルーチンが起動されます。",
			},
			{
				ID:          "6-1-2",
				Text:        "ゴルーチンの初期スタックサイズは？",
				Options:     []string{"1MB", "約2KB", "64KB", "8MB"},
				Answer:      1,
				Explanation: "ゴルーチンの初期スタックサイズは約2KBと非常に軽量です。必要に応じて自動的に拡張されます。",
			},
			{
				ID:          "6-1-3",
				Text:        "メインゴルーチンが終了するとどうなる？",
				Options:     []string{"他のゴルーチンは続行する", "全ゴルーチンが終了する", "デッドロックになる", "エラーが返る"},
				Answer:      1,
				Explanation: "メインゴルーチンが終了すると、プログラム全体が終了し、全ゴルーチンも終了します。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "6-2",
		ChapterID: 6,
		Title:     "チャネル",
		Content: `**チャネル (channel)** はゴルーチン間でデータを安全にやり取りするための仕組みです。

「メモリを共有して通信するのではなく、通信によってメモリを共有せよ」というGoの哲学を体現しています。

チャネルの種類:
- **バッファなしチャネル**: <code>make(chan 型)</code> — 送受信が同期される
- **バッファ付きチャネル**: <code>make(chan 型, サイズ)</code> — バッファが一杯になるまで送信はブロックされない

<code><-</code> 演算子で送受信を行います。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "チャネルの基本",
				Code: `package main

import "fmt"

func main() {
    // バッファなしチャネル
    ch := make(chan string)

    go func() {
        ch <- "Hello from goroutine!"
    }()

    msg := <-ch // 受信（ブロックする）
    fmt.Println(msg)

    // バッファ付きチャネル
    buffered := make(chan int, 3)
    buffered <- 1
    buffered <- 2
    buffered <- 3
    // buffered <- 4 // ブロックする（バッファが一杯）

    fmt.Println(<-buffered) // 1
    fmt.Println(<-buffered) // 2
    fmt.Println(<-buffered) // 3
}`,
			},
			{
				Title: "チャネルでの同期パターン",
				Code: `package main

import "fmt"

func sum(nums []int, ch chan int) {
    total := 0
    for _, n := range nums {
        total += n
    }
    ch <- total
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    ch := make(chan int)

    // 前半と後半を並行で計算
    go sum(nums[:5], ch)
    go sum(nums[5:], ch)

    a, b := <-ch, <-ch
    fmt.Printf("合計: %d + %d = %d\n", a, b, a+b)
}`,
			},
		},
		Notes: []string{
			"バッファなしチャネルは送信と受信が同時に行われるまでブロックします",
			"close(ch) でチャネルを閉じると、受信側は残りのデータを読んだ後ゼロ値を受け取ります",
			"range でチャネルが閉じられるまでループで受信できます",
		},
		Exercise: &models.Exercise{
			Title:       "メッセージの送受信",
			Description: "string型のチャネルを作成し、ゴルーチンから \"Ping\" という文字列を送信し、メイン関数で受信して表示してください。",
			StarterCode: `package main

import "fmt"

func main() {
    // チャネル作成
    message := make(chan string)
    
    // ゴルーチン起動
    go func() {
        // メッセージ送信
        
    }()
    
    // 受信して表示
    fmt.Println(<-message)
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "6-2",
		Questions: []models.Question{
			{
				ID:          "6-2-1",
				Text:        "バッファなしチャネルの特徴は？",
				Options:     []string{"データを蓄積できる", "送受信が同期される", "複数の値を保持できる", "一方向のみ"},
				Answer:      1,
				Explanation: "バッファなしチャネルは送信側と受信側が同時に準備できるまで両方がブロックされます。",
			},
			{
				ID:          "6-2-2",
				Text:        "ch <- value の意味は？",
				Options:     []string{"チャネルから受信", "チャネルに送信", "チャネルを閉じる", "チャネルの長さ"},
				Answer:      1,
				Explanation: "ch <- value はチャネル ch に value を送信します。<-ch で受信します。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "6-3",
		ChapterID: 6,
		Title:     "select文",
		Content: `<code>select</code> 文は複数のチャネル操作を同時に待機する仕組みです。

switch文に似ていますが、各caseがチャネル操作になっています。複数のcaseが準備完了の場合、ランダムに1つが選ばれます。

<code>default</code> ケースを使うとノンブロッキング操作ができます。<code>time.After()</code> と組み合わせてタイムアウトを実装することもよくあります。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "selectの基本とタイムアウト",
				Code: `package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(100 * time.Millisecond)
        ch1 <- "ch1のデータ"
    }()

    go func() {
        time.Sleep(200 * time.Millisecond)
        ch2 <- "ch2のデータ"
    }()

    // 2つのチャネルから受信
    for i := 0; i < 2; i++ {
        select {
        case msg := <-ch1:
            fmt.Println("受信:", msg)
        case msg := <-ch2:
            fmt.Println("受信:", msg)
        case <-time.After(1 * time.Second):
            fmt.Println("タイムアウト")
        }
    }
}`,
			},
		},
		Notes: []string{
			"複数のcaseが同時に準備完了の場合、ランダムに選ばれます",
			"default ケースを入れるとノンブロッキングになります",
			"for + select はイベントループのパターンでよく使われます",
		},
		Exercise: &models.Exercise{
			Title:       "早い者勝ち",
			Description: "2つのチャネルを作成し、それぞれ異なる時間待機してからデータを送信するゴルーチンを起動します。selectを使って、先に到着したデータのみを表示してください（1回だけ受信）。",
			StarterCode: `package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "亀"
    }()
    
    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "ウサギ"
    }()
    
    // selectで待機（先に到着した方だけ表示）
    select {
    
    
    }
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "6-3",
		Questions: []models.Question{
			{
				ID:          "6-3-1",
				Text:        "selectで複数のcaseが同時に準備完了の場合は？",
				Options:     []string{"最初のcaseが実行される", "全caseが実行される", "ランダムに1つが選ばれる", "エラーになる"},
				Answer:      2,
				Explanation: "selectでは複数のcaseが準備完了の場合、ランダムに1つが選ばれます。公平性を保つための設計です。",
			},
			{
				ID:          "6-3-2",
				Text:        "selectにdefaultケースを追加するとどうなる？",
				Options:     []string{"常にdefaultが実行される", "ノンブロッキングになる", "エラーハンドリングになる", "無限ループになる"},
				Answer:      1,
				Explanation: "defaultケースがあると、どのチャネルも準備完了でない場合にdefaultが即座に実行されます。",
			},
		},
	})

	s.addLesson(models.Lesson{
		ID:        "6-4",
		ChapterID: 6,
		Title:     "syncパッケージ",
		Content: `<code>sync</code> パッケージは低レベルの同期プリミティブを提供します。

主要な型:
- <code>sync.WaitGroup</code>: 複数のゴルーチンの完了を待つ
- <code>sync.Mutex</code>: 排他制御（ミューテックス）
- <code>sync.RWMutex</code>: 読み書きロック
- <code>sync.Once</code>: 一度だけ実行を保証

チャネルよりも単純な同期が必要な場合に使います。`,
		CodeExamples: []models.CodeExample{
			{
				Title: "WaitGroupとMutex",
				Code: `package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    var mu sync.Mutex
    counter := 0

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }

    wg.Wait()
    fmt.Println("カウンター:", counter) // 1000
}`,
			},
			{
				Title: "sync.Once",
				Code: `package main

import (
    "fmt"
    "sync"
)

var once sync.Once

func initialize() {
    fmt.Println("初期化は一度だけ実行されます")
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            once.Do(initialize) // 最初の1回だけ実行
        }()
    }
    wg.Wait()
}`,
			},
		},
		Notes: []string{
			"WaitGroup.Add() は必ずゴルーチン起動前に呼びます",
			"defer wg.Done() のパターンで確実にカウントを減らします",
			"チャネルで解決できる場合はチャネルを優先しましょう",
		},
		Exercise: &models.Exercise{
			Title:       "WaitGroupでの待機",
			Description: "sync.WaitGroupを使って、3つのゴルーチンの完了を待機するプログラムを完成させてください。各ゴルーチンは単に \"Done!\" と表示するだけで構いません。",
			StarterCode: `package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 3; i++ {
        // Add呼び出し
        
        go func(id int) {
            // Done呼び出し
            
            fmt.Printf("Goroutine %d finished\n", id)
        }(i)
    }
    
    // 待機
    
    fmt.Println("All done")
}`,
		},
	})

	s.addQuiz(models.Quiz{
		LessonID: "6-4",
		Questions: []models.Question{
			{
				ID:          "6-4-1",
				Text:        "sync.WaitGroup の用途は？",
				Options:     []string{"データを共有する", "ゴルーチンの完了を待つ", "ゴルーチンを作成する", "チャネルを制御する"},
				Answer:      1,
				Explanation: "WaitGroupは複数のゴルーチンの完了を待つために使います。Add, Done, Wait の3メソッドで構成されます。",
			},
			{
				ID:          "6-4-2",
				Text:        "sync.Once の特徴は？",
				Options:     []string{"毎回実行する", "指定回数だけ実行する", "一度だけ実行を保証する", "並行実行を許可する"},
				Answer:      2,
				Explanation: "sync.Once は Do メソッドに渡された関数を、複数のゴルーチンから呼ばれても一度だけ実行することを保証します。",
			},
		},
	})
}
