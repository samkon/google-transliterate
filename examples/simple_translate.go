//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

var content = "كامل الحارونى متفرع من احمد فخرى امام حديقه الطفل مدينه نصر القاهره الدور السابع شقه ١٦"

func main() {
	c := translator.Config{
		Proxy:       "http://127.0.0.1:7890",
		UserAgent:   []string{"Custom Agent"},
		ServiceUrls: []string{"translate.google.com.hk"},
	}
	t := translator.New(c)
	result, err := t.Transliterate(content, "auto", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Text)
}
