package hangul

import (
	"fmt"
)

func ExampleSyllable() {
	var s string
	s = "읽기읽기" // %u110B%u1175%u11B0%u1100%u1175 + %uC77D%uAE30
	syllables, err := Parse(s)
	if err != nil {
		panic(err)
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
	// [[ㅇ ㅣ ㄺ] [ㄱ ㅣ] [ㅇ ㅣ ㄺ] [ㄱ ㅣ]]
	// [[ㅇ ㅣ ㄹ ㄱ] [ㄱ ㅣ] [ㅇ ㅣ ㄹ ㄱ] [ㄱ ㅣ]]
}
