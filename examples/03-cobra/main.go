package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// グローバルフラグ 🌍
	verbose bool
)

func main() {
	// ルートコマンドの作成 🌳
	var rootCmd = &cobra.Command{
		Use:   "mycli",
		Short: "めっちゃイケてる CLI ツール ✨",
		Long:  `Cobra を使ったサブコマンド対応 CLI ツールだよ〜！`,
	}

	// グローバルフラグの追加
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "詳細表示モード 👀")

	// greet サブコマンド 👋
	var greetCmd = &cobra.Command{
		Use:   "greet [name]",
		Short: "挨拶するコマンドだよ 💕",
		Long:  `指定された名前に挨拶するよ〜！名前を省略するとデフォルトで挨拶するよ✨`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := "ギャルエンジニア"
			if len(args) > 0 {
				name = args[0]
			}

			emoji, _ := cmd.Flags().GetString("emoji")
			times, _ := cmd.Flags().GetInt("times")

			if verbose {
				fmt.Printf("[DEBUG] 名前: %s, 絵文字: %s, 回数: %d\n", name, emoji, times)
			}

			for i := 0; i < times; i++ {
				fmt.Printf("Hello, %s! %s\n", name, emoji)
			}
		},
	}

	// greet コマンド専用のフラグ 🚩
	greetCmd.Flags().StringP("emoji", "e", "✨", "使う絵文字を指定 🎨")
	greetCmd.Flags().IntP("times", "t", 1, "繰り返し回数 🔢")

	// info サブコマンド ℹ️
	var infoCmd = &cobra.Command{
		Use:   "info",
		Short: "システム情報を表示 📊",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("=== システム情報 ===")
			fmt.Printf("OS: %s\n", "Go Runtime")
			fmt.Printf("CLI ツール: mycli v1.0.0\n")
			fmt.Printf("Author: ギャルエンジニア 💕\n")
			fmt.Println("===================")

			if verbose {
				fmt.Println("\n[DEBUG] 詳細モードで実行中 👀")
			}
		},
	}

	// サブコマンドを追加 ➕
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(infoCmd)

	// コマンドを実行！🔥
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}
}
