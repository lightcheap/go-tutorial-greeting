package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// ログエントリプレフィックスや時間、ソースファイル、行番号の印刷を無効にするフラグなど、
	// 事前定義されたロガーのプロパティを設定します…
	// って書いてあったけどここ意味がわからん
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// 名前のスライスを作成
	names := []string{"まさよし", "やすよ", "あいか"}

	// グリーティングのメッセージをリクエストする。
	// 名前にはスライスを指定
	messages, err := greetings.Hellos(names)

	// エラーが返ってきていたら、エラーメッセージを表示して処理を終了する
	if err != nil {
		log.Fatal(err)
	}

	// エラーが返されていないばあい、
	// グリーティングメッセージを取得して出力。
	fmt.Println(messages)
}
