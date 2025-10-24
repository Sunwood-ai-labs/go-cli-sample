package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person 構造体 👤
type Person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Email    string   `json:"email"`
	Hobbies  []string `json:"hobbies"`
	IsActive bool     `json:"is_active"`
}

// Config 構造体（ネストした例）⚙️
type Config struct {
	AppName string `json:"app_name"`
	Version string `json:"version"`
	Admin   Person `json:"admin"`
}

func main() {
	fmt.Println("📋 JSON パース・生成サンプル ✨")
	fmt.Println("========================================")
	fmt.Println()

	// 1. JSON 文字列を Go 構造体にデコード 🔄
	fmt.Println("1️⃣ JSON をデコードするよ〜")
	jsonStr := `{
		"name": "ギャルエンジニア",
		"age": 25,
		"email": "gal@example.com",
		"hobbies": ["コーディング", "カフェ巡り", "ショッピング"],
		"is_active": true
	}`

	var person Person
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		fmt.Printf("❌ JSON デコード失敗: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ デコード成功！")
	fmt.Printf("名前: %s\n", person.Name)
	fmt.Printf("年齢: %d\n", person.Age)
	fmt.Printf("メール: %s\n", person.Email)
	fmt.Printf("趣味: %v\n", person.Hobbies)
	fmt.Printf("アクティブ: %v\n", person.IsActive)
	fmt.Println()

	// 2. Go 構造体を JSON にエンコード 🔄
	fmt.Println("2️⃣ 構造体を JSON にエンコードするよ〜")
	newPerson := Person{
		Name:     "イケてるエンジニア",
		Age:      30,
		Email:    "cool@example.com",
		Hobbies:  []string{"Go言語", "OSS開発", "技術ブログ"},
		IsActive: true,
	}

	// コンパクトな JSON（1行）
	jsonData, err := json.Marshal(newPerson)
	if err != nil {
		fmt.Printf("❌ JSON エンコード失敗: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("📄 コンパクトな JSON:")
	fmt.Println(string(jsonData))
	fmt.Println()

	// 3. きれいにフォーマットされた JSON を生成 ✨
	fmt.Println("3️⃣ きれいにフォーマットされた JSON を生成するよ〜")
	prettyJSON, err := json.MarshalIndent(newPerson, "", "  ")
	if err != nil {
		fmt.Printf("❌ JSON エンコード失敗: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("📄 フォーマット済み JSON:")
	fmt.Println(string(prettyJSON))
	fmt.Println()

	// 4. ネストした構造体の JSON 処理 🏗️
	fmt.Println("4️⃣ ネストした構造体を処理するよ〜")
	config := Config{
		AppName: "Go CLI Sample",
		Version: "1.0.0",
		Admin: Person{
			Name:     "管理者ちゃん",
			Age:      28,
			Email:    "admin@example.com",
			Hobbies:  []string{"サーバー管理", "セキュリティ"},
			IsActive: true,
		},
	}

	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Printf("❌ JSON エンコード失敗: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("📄 ネストした JSON:")
	fmt.Println(string(configJSON))
	fmt.Println()

	// 5. JSON ファイルに保存 💾
	fmt.Println("5️⃣ JSON ファイルに保存するよ〜")
	filename := "config.json"
	err = os.WriteFile(filename, configJSON, 0644)
	if err != nil {
		fmt.Printf("❌ ファイル書き込み失敗: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✅ %s に保存したよ！\n", filename)
	fmt.Println()

	// 6. JSON ファイルから読み込み 📖
	fmt.Println("6️⃣ JSON ファイルから読み込むよ〜")
	fileData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("❌ ファイル読み込み失敗: %v\n", err)
		os.Exit(1)
	}

	var loadedConfig Config
	err = json.Unmarshal(fileData, &loadedConfig)
	if err != nil {
		fmt.Printf("❌ JSON デコード失敗: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ ファイルから読み込み成功！")
	fmt.Printf("アプリ名: %s\n", loadedConfig.AppName)
	fmt.Printf("バージョン: %s\n", loadedConfig.Version)
	fmt.Printf("管理者: %s (%s)\n", loadedConfig.Admin.Name, loadedConfig.Admin.Email)
	fmt.Println()

	// 7. クリーンアップ 🧹
	fmt.Println("7️⃣ クリーンアップするよ〜")
	err = os.Remove(filename)
	if err != nil {
		fmt.Printf("⚠️ ファイル削除失敗: %v\n", err)
	} else {
		fmt.Printf("✅ %s を削除したよ！\n", filename)
	}
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("すべての JSON 操作が完了したよ〜！🎉")
}
