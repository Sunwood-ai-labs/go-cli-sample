package main

import (
	"fmt"
	"os"
)

func printBanner() {
	banner := `
   ____           ____ _     ___   ____                        _
  / ___| ___     / ___| |   |_ _| / ___|  __ _ _ __ ___  _ __ | | ___
 | |  _ / _ \   | |   | |    | |  \___ \ / _' | '_ ' _ \| '_ \| |/ _ \
 | |_| | (_) |  | |___| |___ | |   ___) | (_| | | | | | | |_) | |  __/
  \____|\___/    \____|_____|___| |____/ \__,_|_| |_| |_| .__/|_|\___|
                                                         |_|
                        ✨ シンプルでイケてるCLIツール ✨
	`
	fmt.Println(banner)
}

func main() {
	printBanner()
	fmt.Println("\nHello, ギャルエンジニア from Go CLI! ✨")
	if len(os.Args) > 1 {
		fmt.Printf("引数うけとったよ〜: %s\n", os.Args[1])
	}
}
