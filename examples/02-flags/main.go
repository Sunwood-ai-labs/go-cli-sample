package main

import (
	"flag"
	"fmt"
)

func main() {
	// フラグの定義だよ〜！🚩
	name := flag.String("name", "ギャルエンジニア", "あなたの名前を指定してね💕")
	count := flag.Int("count", 1, "何回繰り返すか指定だよ🔢")
	verbose := flag.Bool("verbose", false, "詳細表示モード👀")
	emoji := flag.String("emoji", "✨", "使う絵文字を指定してね🎨")

	// フラグをパース（これ大事🔥）
	flag.Parse()

	// verbose モードの場合は設定値を表示
	if *verbose {
		fmt.Println("=== 設定内容 ===")
		fmt.Printf("名前: %s\n", *name)
		fmt.Printf("繰り返し回数: %d\n", *count)
		fmt.Printf("絵文字: %s\n", *emoji)
		fmt.Println("================")
		fmt.Println()
	}

	// メイン処理✨
	for i := 0; i < *count; i++ {
		fmt.Printf("Hello, %s! %s\n", *name, *emoji)
	}

	// 追加の引数がある場合は表示
	if len(flag.Args()) > 0 {
		fmt.Printf("\nフラグ以外の引数: %v\n", flag.Args())
	}
}
