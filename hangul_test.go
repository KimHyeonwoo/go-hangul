package hangul

import (
	"fmt"
)

func Example() {
	var s string
	s = "읽기읽기" // %u110B%u1175%u11B0%u1100%u1175 + %uC77D%uAE30
	syllables, err := Parse(s)
	if err != nil {
		panic(err)
	}

	// Result for DefaultDecompose
	fmt.Printf("%c\n", syllables.Decompose(true, DefaultDecompose))

	// Result for FullDecompose
	fmt.Printf("%c\n", syllables.Decompose(true, FullDecompose))

	// Output:
	// [[ㅇ ㅣ ㄺ] [ㄱ ㅣ] [ㅇ ㅣ ㄺ] [ㄱ ㅣ]]
	// [[ㅇ ㅣ ㄹ ㄱ] [ㄱ ㅣ] [ㅇ ㅣ ㄹ ㄱ] [ㄱ ㅣ]]
}
