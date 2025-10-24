package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("📁 ファイル操作サンプル ✨")
	fmt.Println("========================================")
	fmt.Println()

	filename := "sample.txt"
	content := "Hello, ギャルエンジニア！✨\nGo でファイル操作できるよ〜🔥\n"

	// 1. ファイルの作成と書き込み 📝
	fmt.Println("1️⃣ ファイルを作成して書き込むよ〜")
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("❌ エラー: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✅ %s を作成したよ！\n\n", filename)

	// 2. ファイルの存在確認 🔍
	fmt.Println("2️⃣ ファイルが存在するか確認するよ〜")
	if _, err := os.Stat(filename); err == nil {
		fmt.Printf("✅ %s は存在するよ！\n\n", filename)
	} else if os.IsNotExist(err) {
		fmt.Printf("❌ %s は存在しないよ😢\n\n", filename)
	} else {
		fmt.Printf("⚠️ エラー: %v\n\n", err)
	}

	// 3. ファイルの読み込み 📖
	fmt.Println("3️⃣ ファイルの内容を読み込むよ〜")
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("❌ エラー: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("📄 ファイルの内容:")
	fmt.Println("----------------------------------------")
	fmt.Print(string(data))
	fmt.Println("----------------------------------------")
	fmt.Println()

	// 4. ファイル情報を取得 📊
	fmt.Println("4️⃣ ファイル情報を取得するよ〜")
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("❌ エラー: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("ファイル名: %s\n", fileInfo.Name())
	fmt.Printf("ファイルサイズ: %d bytes\n", fileInfo.Size())
	fmt.Printf("パーミッション: %s\n", fileInfo.Mode())
	fmt.Printf("最終更新時刻: %s\n", fileInfo.ModTime())
	fmt.Println()

	// 5. ファイルに追記 ➕
	fmt.Println("5️⃣ ファイルに追記するよ〜")
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("❌ エラー: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	appendContent := "追記した内容だよ〜💕\n"
	if _, err := f.WriteString(appendContent); err != nil {
		fmt.Printf("❌ エラー: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ ファイルに追記したよ！")
	fmt.Println()

	// 6. 追記後の内容を確認 👀
	fmt.Println("6️⃣ 追記後の内容を確認するよ〜")
	data, err = os.ReadFile(filename)
	if err != nil {
		fmt.Printf("❌ エラー: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("📄 ファイルの内容:")
	fmt.Println("----------------------------------------")
	fmt.Print(string(data))
	fmt.Println("----------------------------------------")
	fmt.Println()

	// 7. ファイルの削除 🗑️
	fmt.Println("7️⃣ ファイルを削除するよ〜")
	err = os.Remove(filename)
	if err != nil {
		fmt.Printf("❌ エラー: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✅ %s を削除したよ！\n", filename)
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("すべての操作が完了したよ〜！🎉")
}
