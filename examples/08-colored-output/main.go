package main

import (
	"fmt"
)

// ANSI カラーコード 🎨
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"

	// 背景色 🎭
	BgRed    = "\033[41m"
	BgGreen  = "\033[42m"
	BgYellow = "\033[43m"
	BgBlue   = "\033[44m"
	BgPurple = "\033[45m"
	BgCyan   = "\033[46m"

	// スタイル ✨
	Bold      = "\033[1m"
	Underline = "\033[4m"
	Blink     = "\033[5m"
)

// カラー出力のヘルパー関数 🎨
func colorize(color, text string) string {
	return color + text + ColorReset
}

// 成功メッセージ ✅
func success(message string) {
	fmt.Println(colorize(ColorGreen, "✅ "+message))
}

// エラーメッセージ ❌
func errorMsg(message string) {
	fmt.Println(colorize(ColorRed, "❌ "+message))
}

// 警告メッセージ ⚠️
func warning(message string) {
	fmt.Println(colorize(ColorYellow, "⚠️  "+message))
}

// 情報メッセージ ℹ️
func info(message string) {
	fmt.Println(colorize(ColorBlue, "ℹ️  "+message))
}

func main() {
	fmt.Println("🌈 カラフルな出力サンプル ✨")
	fmt.Println("========================================")
	fmt.Println()

	// 1. 基本的な色 🎨
	fmt.Println("1️⃣ 基本的な色を表示するよ〜")
	fmt.Println(colorize(ColorRed, "赤色のテキスト"))
	fmt.Println(colorize(ColorGreen, "緑色のテキスト"))
	fmt.Println(colorize(ColorYellow, "黄色のテキスト"))
	fmt.Println(colorize(ColorBlue, "青色のテキスト"))
	fmt.Println(colorize(ColorPurple, "紫色のテキスト"))
	fmt.Println(colorize(ColorCyan, "シアン色のテキスト"))
	fmt.Println(colorize(ColorWhite, "白色のテキスト"))
	fmt.Println()

	// 2. スタイル付きテキスト ✨
	fmt.Println("2️⃣ スタイルを適用するよ〜")
	fmt.Println(colorize(Bold, "太字のテキスト"))
	fmt.Println(colorize(Underline, "下線付きテキスト"))
	fmt.Println(colorize(Bold+ColorRed, "太字の赤色テキスト"))
	fmt.Println(colorize(Bold+Underline+ColorBlue, "太字＋下線＋青色"))
	fmt.Println()

	// 3. 背景色 🎭
	fmt.Println("3️⃣ 背景色を設定するよ〜")
	fmt.Println(colorize(BgRed+ColorWhite, " 赤い背景 "))
	fmt.Println(colorize(BgGreen+ColorWhite, " 緑の背景 "))
	fmt.Println(colorize(BgYellow+ColorWhite, " 黄色の背景 "))
	fmt.Println(colorize(BgBlue+ColorWhite, " 青い背景 "))
	fmt.Println(colorize(BgPurple+ColorWhite, " 紫の背景 "))
	fmt.Println(colorize(BgCyan+ColorWhite, " シアンの背景 "))
	fmt.Println()

	// 4. 実用的な使い方 💼
	fmt.Println("4️⃣ 実用的なメッセージ表示だよ〜")
	success("処理が成功したよ！")
	errorMsg("エラーが発生したよ😢")
	warning("これは警告メッセージだよ！")
	info("お知らせがあるよ〜")
	fmt.Println()

	// 5. 組み合わせ例 🎨
	fmt.Println("5️⃣ いろんな組み合わせ例だよ〜")
	fmt.Printf("%s %s %s\n",
		colorize(ColorRed, "エラー:"),
		"ファイルが見つからないよ😢",
		colorize(ColorYellow, "(code: 404)"))

	fmt.Printf("%s %s %s\n",
		colorize(Bold+ColorGreen, "成功!"),
		"ビルドが完了したよ〜",
		colorize(ColorCyan, "(0.5秒)"))

	fmt.Printf("%s %s\n",
		colorize(Bold+ColorBlue, "📊 進捗:"),
		colorize(ColorGreen, "████████") + colorize(ColorWhite, "░░") + " 80%")
	fmt.Println()

	// 6. レインボー表示 🌈
	fmt.Println("6️⃣ レインボー表示だよ〜")
	rainbow := []string{ColorRed, ColorYellow, ColorGreen, ColorCyan, ColorBlue, ColorPurple}
	text := "カラフル！"
	for i, char := range text {
		color := rainbow[i%len(rainbow)]
		fmt.Print(colorize(color, string(char)))
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println(colorize(Bold+ColorGreen, "すべてのカラー表示が完了したよ〜！🎉"))
}
