package greetings

import (
	"regexp"
	"testing"
)

// 名前を使用して
func TestHelloName(t *testing.T) {
	name := "mike" // ここの名前、ひらがなだとテストがパスしない。。
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("mike")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("mike") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// テストを試すところから
