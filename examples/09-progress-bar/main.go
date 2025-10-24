package main

import (
	"fmt"
	"time"
)

// カラーコード 🎨
const (
	ColorReset = "\033[0m"
	ColorGreen = "\033[32m"
	ColorCyan  = "\033[36m"
	ColorWhite = "\033[37m"
)

// シンプルなプログレスバーを表示 📊
func simpleProgressBar(percent int) {
	barLength := 50
	filled := percent * barLength / 100
	empty := barLength - filled

	bar := ""
	for i := 0; i < filled; i++ {
		bar += "█"
	}
	for i := 0; i < empty; i++ {
		bar += "░"
	}

	// \r でカーソルを行頭に戻して上書き 🔄
	fmt.Printf("\r[%s] %d%%", bar, percent)
}

// カラフルなプログレスバーを表示 🌈
func colorfulProgressBar(percent int) {
	barLength := 50
	filled := percent * barLength / 100
	empty := barLength - filled

	bar := ColorGreen
	for i := 0; i < filled; i++ {
		bar += "█"
	}
	bar += ColorWhite
	for i := 0; i < empty; i++ {
		bar += "░"
	}
	bar += ColorReset

	// 進捗に応じてメッセージを変える 💬
	var status string
	switch {
	case percent < 25:
		status = "開始中..."
	case percent < 50:
		status = "進行中..."
	case percent < 75:
		status = "もうちょい！"
	case percent < 100:
		status = "あと少し！"
	default:
		status = "完了！✨"
	}

	fmt.Printf("\r[%s] %d%% - %s", bar, percent, status)
}

// スピナーアニメーション 🔄
func spinner(message string, duration time.Duration) {
	spinChars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	start := time.Now()

	i := 0
	for time.Since(start) < duration {
		fmt.Printf("\r%s %s", spinChars[i%len(spinChars)], message)
		i++
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("\r✅ %s - 完了！\n", message)
}

// ダウンロード風プログレスバー 📥
func downloadProgressBar(current, total int) {
	percent := current * 100 / total
	barLength := 30
	filled := percent * barLength / 100
	empty := barLength - filled

	bar := ColorCyan
	for i := 0; i < filled; i++ {
		bar += "="
	}
	if filled < barLength {
		bar += ">"
		filled++
	}
	bar += ColorWhite
	for i := filled; i < barLength; i++ {
		bar += " "
	}
	bar += ColorReset

	// サイズを MB で表示
	currentMB := float64(current) / 1024.0
	totalMB := float64(total) / 1024.0

	fmt.Printf("\rダウンロード中: [%s] %d%% (%.1f/%.1f MB)", bar, percent, currentMB, totalMB)
}

func main() {
	fmt.Println("⏳ プログレスバーサンプル ✨")
	fmt.Println("========================================")
	fmt.Println()

	// 1. シンプルなプログレスバー 📊
	fmt.Println("1️⃣ シンプルなプログレスバー")
	for i := 0; i <= 100; i++ {
		simpleProgressBar(i)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println() // 改行
	fmt.Println()

	// 2. カラフルなプログレスバー 🌈
	fmt.Println("2️⃣ カラフルなプログレスバー")
	for i := 0; i <= 100; i++ {
		colorfulProgressBar(i)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println() // 改行
	fmt.Println()

	// 3. スピナーアニメーション 🔄
	fmt.Println("3️⃣ スピナーアニメーション")
	spinner("データを処理中", 2*time.Second)
	spinner("ファイルを読み込み中", 2*time.Second)
	spinner("サーバーに接続中", 2*time.Second)
	fmt.Println()

	// 4. ダウンロード風プログレスバー 📥
	fmt.Println("4️⃣ ダウンロード風プログレスバー")
	total := 10240 // 10 MB
	for i := 0; i <= total; i += 128 {
		downloadProgressBar(i, total)
		time.Sleep(20 * time.Millisecond)
	}
	downloadProgressBar(total, total) // 最後に 100% を表示
	fmt.Println() // 改行
	fmt.Println()

	// 5. マルチステッププログレス 📝
	fmt.Println("5️⃣ マルチステッププログレス")
	steps := []string{
		"依存関係をインストール中",
		"ソースコードをコンパイル中",
		"テストを実行中",
		"ビルド成果物を作成中",
		"デプロイ中",
	}

	for i, step := range steps {
		spinner(step, 1*time.Second)
	}
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println(ColorGreen + "すべての処理が完了したよ〜！🎉" + ColorReset)
}
