package wjjutils

import (
	"fmt"
	"github.com/lyyzwjj/kana"
)

func A() {
	fmt.Println(kana.NormalizeRomaji("買う"))
	fmt.Println(kana.NormalizeRomaji("Katakana"))
}
func ChangeVerb(Kara, Kanji *string) {
	fmt.Println(kana.NormalizeRomaji("買う"))
}
