package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, ギャルエンジニア from Go CLI! ✨")
	if len(os.Args) > 1 {
		fmt.Printf("引数うけとったよ〜: %s\n", os.Args[1])
	}
}
