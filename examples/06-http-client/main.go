package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// JSONPlaceholder の User 構造体 👤
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
}

func main() {
	fmt.Println("🌐 HTTP クライアントサンプル ✨")
	fmt.Println("========================================")
	fmt.Println()

	// 1. シンプルな GET リクエスト 📡
	fmt.Println("1️⃣ シンプルな GET リクエストを送るよ〜")
	url := "https://jsonplaceholder.typicode.com/users/1"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("❌ リクエスト失敗: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Printf("✅ ステータスコード: %d %s\n", resp.StatusCode, resp.Status)
	fmt.Println()

	// 2. レスポンスボディを読み取る 📖
	fmt.Println("2️⃣ レスポンスを読み取るよ〜")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ レスポンス読み取り失敗: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("📄 レスポンスボディ:")
	fmt.Println("----------------------------------------")
	fmt.Println(string(body))
	fmt.Println("----------------------------------------")
	fmt.Println()

	// 3. JSON をパースして構造体に変換 🔄
	fmt.Println("3️⃣ JSON をパースするよ〜")
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Printf("❌ JSON パース失敗: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("👤 ユーザー情報:")
	fmt.Printf("  ID: %d\n", user.ID)
	fmt.Printf("  名前: %s\n", user.Name)
	fmt.Printf("  ユーザー名: %s\n", user.Username)
	fmt.Printf("  メール: %s\n", user.Email)
	fmt.Printf("  電話: %s\n", user.Phone)
	fmt.Printf("  ウェブサイト: %s\n", user.Website)
	fmt.Println()

	// 4. カスタム HTTP クライアントでタイムアウト設定 ⏱️
	fmt.Println("4️⃣ タイムアウト付きリクエストを送るよ〜")
	client := &http.Client{
		Timeout: 10 * time.Second, // 10 秒でタイムアウト
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("❌ リクエスト作成失敗: %v\n", err)
		os.Exit(1)
	}

	// ヘッダーを追加 📋
	req.Header.Set("User-Agent", "Go-CLI-Sample/1.0")
	req.Header.Set("Accept", "application/json")

	resp2, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ リクエスト失敗: %v\n", err)
		os.Exit(1)
	}
	defer resp2.Body.Close()

	fmt.Printf("✅ ステータスコード: %d\n", resp2.StatusCode)
	fmt.Printf("✅ Content-Type: %s\n", resp2.Header.Get("Content-Type"))
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("すべてのリクエストが完了したよ〜！🎉")
}
