package hangul

const (
	DefaultDecompose = iota
	FullDecompose
)

// Syllable represents a syllable in Korean.
// It consists of a Choseong (initial consonant), a Jungseong (medial vowel), and a Jongseong (final consonant).
// Unicode `Hangul Syllables` range: U+AC00 ~ U+D7A3
// Unicode for Syllable is [(index for initial consonant) × 588 + (index for medial vowel) × 28 + (index for final consonant)] + 0xAC00
// where 588 = 21 × 28, 21 is the number of medial vowels (JungseongCount), and 28 is the number of final consonants (JongseongCount).
// ref: https://en.wikipedia.org/wiki/Korean_language_and_computers#Hangul_in_Unicode
type Syllable struct {
	Choseong  Choseong
	Jungseong Jungseong
	Jongseong Jongseong
}

const SyllableBase = 0xAC00

// Rune returns the rune of the `Hangul Syllables`.
func (s Syllable) Rune() rune {
	return rune(SyllableBase + int(s.Choseong)*JungseongCount*JongseongCount + int(s.Jungseong)*JongseongCount + int(s.Jongseong))
}

// Decompose returns the `decomposed` slice of runes for `Hangul Jamo` or `Hangul Compatibility Jamo`.
// If `compatibility` is true, it returns the slice of runes for `Hangul Compatibility Jamo`.
// Otherwise, it returns the slice of runes for `Hangul Jamo`.
// If `level` is FullDecompose, it decomposes a double consonants and double vowels into two to three runes.
func (s Syllable) Decompose(compatibility bool, level int) []rune {
	if level < DefaultDecompose || level > FullDecompose {
		return nil
	}

	return append(append(s.Choseong.Decompose(compatibility, level), s.Jungseong.Decompose(compatibility, level)...), s.Jongseong.Decompose(compatibility, level)...)
}

// newSyllable receives rune in range of `Hangul Syllables` and returns a Syllable.
func newSyllable(r rune) Syllable {
	cho := (r - SyllableBase) / (JungseongCount * JongseongCount)
	jung := ((r - SyllableBase) % (JungseongCount * JongseongCount)) / JongseongCount
	jong := ((r - SyllableBase) % (JungseongCount * JongseongCount)) % JongseongCount

	return Syllable{Choseong(cho), Jungseong(jung), Jongseong(jong)}
}

type Syllables []Syllable

func (s Syllables) Len() int {
	return len(s)
}

func (s Syllables) String() string {
	runes := make([]rune, 0, len(s))
	for _, syllable := range s {
		runes = append(runes, syllable.Rune())
	}
	return string(runes)
}

func (s Syllables) Decompose(compatibility bool, level int) [][]rune {
	runes := make([][]rune, 0, len(s))
	for _, syllable := range s {
		runes = append(runes, syllable.Decompose(compatibility, level))
	}
	return runes
}
