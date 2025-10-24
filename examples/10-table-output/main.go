package main

import (
	"fmt"
	"strings"
)

// カラーコード 🎨
const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
	Bold        = "\033[1m"
)

// データ構造 📊
type User struct {
	ID     int
	Name   string
	Email  string
	Status string
}

// シンプルなテーブルを表示 📋
func printSimpleTable(users []User) {
	// ヘッダー
	fmt.Println("ID\tName\t\t\tEmail\t\t\t\tStatus")
	fmt.Println(strings.Repeat("-", 80))

	// データ行
	for _, user := range users {
		fmt.Printf("%d\t%-20s\t%-30s\t%s\n", user.ID, user.Name, user.Email, user.Status)
	}
}

// 罫線付きテーブルを表示 📦
func printBoxedTable(users []User) {
	// 列幅を定義
	idWidth := 6
	nameWidth := 25
	emailWidth := 30
	statusWidth := 10

	// 上部の罫線
	fmt.Printf("┌%s┬%s┬%s┬%s┐\n",
		strings.Repeat("─", idWidth),
		strings.Repeat("─", nameWidth),
		strings.Repeat("─", emailWidth),
		strings.Repeat("─", statusWidth))

	// ヘッダー
	fmt.Printf("│ %-*s │ %-*s │ %-*s │ %-*s │\n",
		idWidth-2, "ID",
		nameWidth-2, "Name",
		emailWidth-2, "Email",
		statusWidth-2, "Status")

	// ヘッダーとデータの区切り線
	fmt.Printf("├%s┼%s┼%s┼%s┤\n",
		strings.Repeat("─", idWidth),
		strings.Repeat("─", nameWidth),
		strings.Repeat("─", emailWidth),
		strings.Repeat("─", statusWidth))

	// データ行
	for _, user := range users {
		name := truncate(user.Name, nameWidth-3)
		email := truncate(user.Email, emailWidth-3)
		status := truncate(user.Status, statusWidth-3)

		fmt.Printf("│ %-*d │ %-*s │ %-*s │ %-*s │\n",
			idWidth-2, user.ID,
			nameWidth-2, name,
			emailWidth-2, email,
			statusWidth-2, status)
	}

	// 下部の罫線
	fmt.Printf("└%s┴%s┴%s┴%s┘\n",
		strings.Repeat("─", idWidth),
		strings.Repeat("─", nameWidth),
		strings.Repeat("─", emailWidth),
		strings.Repeat("─", statusWidth))
}

// カラフルなテーブルを表示 🌈
func printColorfulTable(users []User) {
	// 列幅を定義
	idWidth := 6
	nameWidth := 25
	emailWidth := 30
	statusWidth := 10

	// ヘッダー（太字＋シアン）
	header := Bold + ColorCyan
	fmt.Printf("%s┌%s┬%s┬%s┬%s┐%s\n",
		header,
		strings.Repeat("─", idWidth),
		strings.Repeat("─", nameWidth),
		strings.Repeat("─", emailWidth),
		strings.Repeat("─", statusWidth),
		ColorReset)

	fmt.Printf("%s│ %-*s │ %-*s │ %-*s │ %-*s │%s\n",
		header,
		idWidth-2, "ID",
		nameWidth-2, "Name",
		emailWidth-2, "Email",
		statusWidth-2, "Status",
		ColorReset)

	fmt.Printf("%s├%s┼%s┼%s┼%s┤%s\n",
		header,
		strings.Repeat("─", idWidth),
		strings.Repeat("─", nameWidth),
		strings.Repeat("─", emailWidth),
		strings.Repeat("─", statusWidth),
		ColorReset)

	// データ行（ステータスに応じて色分け）
	for _, user := range users {
		name := truncate(user.Name, nameWidth-3)
		email := truncate(user.Email, emailWidth-3)

		// ステータスの色を決定
		var statusColor string
		switch user.Status {
		case "Active":
			statusColor = ColorGreen
		case "Pending":
			statusColor = ColorYellow
		default:
			statusColor = ColorReset
		}

		fmt.Printf("│ %-*d │ %-*s │ %-*s │ %s%-*s%s │\n",
			idWidth-2, user.ID,
			nameWidth-2, name,
			emailWidth-2, email,
			statusColor, statusWidth-2, user.Status, ColorReset)
	}

	// 下部の罫線
	fmt.Printf("%s└%s┴%s┴%s┴%s┘%s\n",
		header,
		strings.Repeat("─", idWidth),
		strings.Repeat("─", nameWidth),
		strings.Repeat("─", emailWidth),
		strings.Repeat("─", statusWidth),
		ColorReset)
}

// 文字列を指定した長さに切り詰める ✂️
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func main() {
	fmt.Println("📊 テーブル出力サンプル ✨")
	fmt.Println("========================================")
	fmt.Println()

	// サンプルデータ
	users := []User{
		{1, "ギャルエンジニア", "gal@example.com", "Active"},
		{2, "イケてるデベロッパー", "cool-dev@example.com", "Active"},
		{3, "かわいいプログラマー", "kawaii-prog@example.com", "Pending"},
		{4, "スーパーハッカー", "super-hacker@example.com", "Active"},
		{5, "最強エンジニア", "saikyo-eng@example.com", "Inactive"},
	}

	// 1. シンプルなテーブル 📋
	fmt.Println("1️⃣ シンプルなテーブル")
	printSimpleTable(users)
	fmt.Println()

	// 2. 罫線付きテーブル 📦
	fmt.Println("2️⃣ 罫線付きテーブル")
	printBoxedTable(users)
	fmt.Println()

	// 3. カラフルなテーブル 🌈
	fmt.Println("3️⃣ カラフルなテーブル")
	printColorfulTable(users)
	fmt.Println()

	// 4. サマリー表示 📊
	fmt.Println("4️⃣ サマリー")
	activeCount := 0
	pendingCount := 0
	inactiveCount := 0

	for _, user := range users {
		switch user.Status {
		case "Active":
			activeCount++
		case "Pending":
			pendingCount++
		case "Inactive":
			inactiveCount++
		}
	}

	fmt.Printf("総ユーザー数: %d\n", len(users))
	fmt.Printf("%sアクティブ%s: %d\n", ColorGreen, ColorReset, activeCount)
	fmt.Printf("%sペンディング%s: %d\n", ColorYellow, ColorReset, pendingCount)
	fmt.Printf("非アクティブ: %d\n", inactiveCount)
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println(ColorGreen + "すべてのテーブル表示が完了したよ〜！🎉" + ColorReset)
}
