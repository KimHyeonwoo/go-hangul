package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/KimHyeonwoo/go-hangul"
)

func main() {
	s := "읽기"
	syllables := make([]hangul.Syllable, 0, utf8.RuneCountInString(s))
	for _, r := range s {
		syllables = append(syllables, hangul.NewSyllable(r))
	}

	fmt.Println("Result for DefaultDecompose")
	for _, syllable := range syllables {
		fmt.Printf("%c\n", syllable.Runes(true, hangul.DefaultDecompose))
	}
	fmt.Println()

	fmt.Println("Result for FullDecompose")
	for _, syllable := range syllables {
		fmt.Printf("%c\n", syllable.Runes(true, hangul.FullDecompose))
	}
}
