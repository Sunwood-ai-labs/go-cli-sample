package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("🎀 インタラクティブ CLI ツール ✨")
	fmt.Println("=====================================")
	fmt.Println()

	// 名前を入力してもらう 📝
	fmt.Print("あなたの名前を教えて〜: ")
	scanner.Scan()
	name := scanner.Text()
	if name == "" {
		name = "ギャルエンジニア"
	}

	// 好きな言語を入力してもらう 💻
	fmt.Print("好きなプログラミング言語は？: ")
	scanner.Scan()
	language := scanner.Text()
	if language == "" {
		language = "Go"
	}

	// 経験年数を入力してもらう 📊
	fmt.Print("何年くらい経験ある？（数字で入力してね）: ")
	scanner.Scan()
	experience := scanner.Text()
	if experience == "" {
		experience = "1"
	}

	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("📋 入力内容の確認")
	fmt.Println("=====================================")
	fmt.Printf("名前: %s\n", name)
	fmt.Printf("好きな言語: %s\n", language)
	fmt.Printf("経験年数: %s 年\n", experience)
	fmt.Println()

	// 確認をとる ✅
	fmt.Print("この内容で OK？ (yes/no): ")
	scanner.Scan()
	confirmation := strings.ToLower(strings.TrimSpace(scanner.Text()))

	if confirmation == "yes" || confirmation == "y" {
		fmt.Println()
		fmt.Printf("よっしゃ〜！%s さん、%s 歴 %s 年なんだね！✨\n", name, language, experience)
		fmt.Println("一緒にいいもの作っていこ〜！💪🔥")
	} else {
		fmt.Println()
		fmt.Println("またやり直してね〜！👋")
	}

	// エラーチェック
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "エラーが発生したよ😢: %v\n", err)
		os.Exit(1)
	}
}
