package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 指定された人の挨拶を返す
func Hello(name string) (string, error) {
	// 名前が無い場合、挨拶しない
	if name == "" {
		return name, errors.New("名前がありません")
	}
	// よくわからんけど、ランダムな形式を使用してメッセージを作成します。
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// テストチュートリアルでの指示。差し替えてみる
// これはテストに失敗する。
// func Hello(name string) (string, error) {
// 	// 名前がない場合、エラーメッセージを返す
// 	if name == "" {
// 		return name, errors.New("名前がありません。")
// 	}
// 	// ランダムフォーマットを用いて、メッセージを作成
// 	message := fmt.Sprintf(randomFormat())

// 	return message, nil
// }

// 指定の氏名を、グリーティングメッセージに関連付けるマップを返す
func Hellos(names []string) (map[string]string, error) {
	// 名前をメッセージに関連付けるマップ
	messages := make(map[string]string)

	// hello関数を呼び出して、各名前のメッセージを取得。
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// マップで取得したメッセージを名前に関連付ける
		messages[name] = message
	}

	return messages, nil
}

// 関数で使用される変数の初期値の設定
func init() {
	rand.Seed(time.Now().UnixNano())
}

// 一連のグリーティングメッセージを返す。
// メッセージはランダムに選択される
func randomFormat() string {
	// 定形メッセージのスライス
	// スライス。。Goでの配列みたいなもんらしい。
	formats := []string{
		"やあ、 %v ようこそ！",
		"さよなら　%v ！",
		"はじめまして %v",
	}

	// ランダムに選択された定形メッセージを返す
	// スライスのランダムインデックス
	return formats[rand.Intn(len(formats))]
}
