# go-hangul

go-hangul 은 한글 자모 분리 및 조합을 위한 Go 패키지입니다.

## Features

- 한글 자모 분리

- [NFC / NFD](https://gist.github.com/Pusnow/aa865fa21f9557fa58d691a8b79f8a6d#%EC%A0%95%EA%B7%9C%ED%99%94) 정규화 방식을 모두 지원

## Installation

```bash
go get github.com/KimHyeonwoo/go-hangul
```

## Usage

```go
package main

import (
	"fmt"

	hangul "github.com/KimHyeonwoo/go-hangul"
)

func main() {
	syllables, _ := hangul.Parse("읽기")

	// 초성, 중성, 종성 분리
	// output: [[ㅇ ㅣ ㄺ] [ㄱ ㅣ]]
	fmt.Printf("%c\n", syllables.Decompose(true, hangul.DefaultDecompose))

	// 초성, 중성, 종성 분리 + 겹자음과 겹모음을 모두 분리
	// output: [[ㅇ ㅣ ㄹ ㄱ] [ㄱ ㅣ]]
	fmt.Printf("%c\n", syllables.Decompose(true, hangul.FullDecompose))
}
```
