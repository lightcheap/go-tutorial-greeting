package Greetings

import "fmt"

// 指定された人の挨拶を返す
func Hello(name string) string {
	message := fmt.Sprintf("やあ、%v。こんにちは！", name)

	return message
}
