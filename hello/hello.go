package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// グリーティングメッセージを取得して出力します。
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
