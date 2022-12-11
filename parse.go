package hangul

import "errors"

// runeType represents the type of rune.
// runeTypeChoseong represents an initial consonant rune in range 0x1100 ~ 0x1112.
// runeTypeJungseong represents a medial vowel rune in range 0x1161 ~ 0x1175.
// runeTypeJongseong represents a final consonant rune in range 0x11A8 ~ 0x11C2.
// runeTypeSyllable represents a syllable rune in range 0xAC00 ~ 0xD7A3.
type runeType int

const (
	runeTypeUndefined runeType = iota
	runeTypeChoseong
	runeTypeJungseong
	runeTypeJongseong
	runeTypeSyllable
)

// parseState represents the state of parsing.
// parseStateInitial represents the initial state.
// parseStateChoseong represents the state after parsing an initial consonant.
// parseStateJungseong represents the state after parsing a medial vowel.
type parseState int

const (
	parseStateInitial parseState = iota
	parseStateChoseong
	parseStateJungseong
	parseStateInvalid
)

// getRuneType returns the runeType or the rune.
func getRuneType(r rune) runeType {
	if r >= 0x1100 && r <= 0x1112 {
		return runeTypeChoseong
	} else if r >= 0x1161 && r <= 0x1175 {
		return runeTypeJungseong
	} else if r >= 0x11A8 && r <= 0x11C2 {
		return runeTypeJongseong
	} else if r >= 0xAC00 && r <= 0xD7A3 {
		return runeTypeSyllable
	} else {
		return runeTypeUndefined
	}
}

// Parse parses the string and returns a slice of Syllables.
func Parse(s string) ([]Syllable, error) {
	var syllables []Syllable
	var state = parseStateInitial

	var choseong Choseong
	var jungseong Jungseong
	var jongseong Jongseong

	runes := []rune(s)
	for _, r := range runes {
		rType := getRuneType(r)
		switch state {
		case parseStateInitial:
			switch rType {
			case runeTypeChoseong:
				choseong = Choseong(r - ChoseongBase)
				state = parseStateChoseong
			case runeTypeSyllable:
				syllables = append(syllables, newSyllable(r))
			default:
				state = parseStateInvalid
			}
		case parseStateChoseong:
			switch rType {
			case runeTypeJungseong:
				jungseong = Jungseong(r - JungseongBase)
				state = parseStateJungseong
			default:
				state = parseStateInvalid
			}
		case parseStateJungseong:
			switch rType {
			case runeTypeJongseong:
				jongseong = Jongseong(r - JongseongBase)
				syllables = append(syllables, Syllable{choseong, jungseong, jongseong})
				choseong = 0
				jungseong = 0
				jongseong = 0
				state = parseStateInitial
			case runeTypeSyllable:
				syllables = append(syllables, Syllable{choseong, jungseong, jongseong})
				syllables = append(syllables, newSyllable(r))
				choseong = 0
				jungseong = 0
				jongseong = 0
				state = parseStateInitial
			default:
				state = parseStateInvalid
			}
		}
	}

	if state == parseStateInitial {
		return syllables, nil
	}

	if state == parseStateJungseong {
		syllables = append(syllables, Syllable{choseong, jungseong, jongseong})
		return syllables, nil
	}

	return nil, errors.New("failed to parse")
}
