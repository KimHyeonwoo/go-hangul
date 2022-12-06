package hangul

import (
	"fmt"
	"unicode/utf8"
)

func ExampleSyllable() {
	s := "읽기"
	syllables := make([]Syllable, 0, utf8.RuneCountInString(s))
	for _, r := range s {
		syllables = append(syllables, NewSyllable(r))
	}

	// Result for DefaultDecompose
	var result1 []string
	for _, syllable := range syllables {
		s := fmt.Sprintf("%c", syllable.Runes(true, DefaultDecompose))
		result1 = append(result1, s)
	}
	fmt.Println(result1)

	// Result for FullDecompose
	var result2 []string
	for _, syllable := range syllables {
		s := fmt.Sprintf("%c", syllable.Runes(true, FullDecompose))
		result2 = append(result2, s)
	}
	fmt.Println(result2)

	// Output:
	// [[ㅇ ㅣ ㄺ] [ㄱ ㅣ]]
	// [[ㅇ ㅣ ㄹ ㄱ] [ㄱ ㅣ]]
}
