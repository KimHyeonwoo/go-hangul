package hangul

// Choseong is an initial consonant in Hangul.
// Unicode `Hangul Jamo` range for Choseong: 0x1100 ~ 0x1112
// Unicode `Hangul Compatibility Jamo` range for Choseong: 0x3131 ~ 0x314E
// ref: https://en.wikipedia.org/wiki/Korean_language_and_computers#Hangul_in_Unicode
type Choseong int32

const (
	ChoseongKiyeok      = Choseong(iota) // ㄱ
	ChoseongSsangKiyeok                  // ㄲ
	ChoseongNieun                        // ㄴ
	ChoseongTikeut                       // ㄷ
	ChoseongSsangTikeut                  // ㄸ
	ChoseongRieul                        // ㄹ
	ChoseongMieum                        // ㅁ
	ChoseongPieup                        // ㅂ
	ChoseongSsangPieup                   // ㅃ
	ChoseongSios                         // ㅅ
	ChoseongSsangSios                    // ㅆ
	ChoseongIeung                        // ㅇ
	ChoseongCieuc                        // ㅈ
	ChoseongSsangCieuc                   // ㅉ
	ChoseongChieuch                      // ㅊ
	ChoseongKhieukh                      // ㅋ
	ChoseongThieuth                      // ㅌ
	ChoseongPhieuph                      // ㅍ
	ChoseongHieuh                        // ㅎ

	ChoseongBase  = 0x1100
	ChoseongCount = 19
)

// CompatibilityRune returns the rune for Unicode `Hangul Compatibility Jamo`.
func (c Choseong) CompatibilityRune() rune {
	switch c {
	case ChoseongKiyeok:
		return 'ㄱ'
	case ChoseongSsangKiyeok:
		return 'ㄲ'
	case ChoseongNieun:
		return 'ㄴ'
	case ChoseongTikeut:
		return 'ㄷ'
	case ChoseongSsangTikeut:
		return 'ㄸ'
	case ChoseongRieul:
		return 'ㄹ'
	case ChoseongMieum:
		return 'ㅁ'
	case ChoseongPieup:
		return 'ㅂ'
	case ChoseongSsangPieup:
		return 'ㅃ'
	case ChoseongSios:
		return 'ㅅ'
	case ChoseongSsangSios:
		return 'ㅆ'
	case ChoseongIeung:
		return 'ㅇ'
	case ChoseongCieuc:
		return 'ㅈ'
	case ChoseongSsangCieuc:
		return 'ㅉ'
	case ChoseongChieuch:
		return 'ㅊ'
	case ChoseongKhieukh:
		return 'ㅋ'
	case ChoseongThieuth:
		return 'ㅌ'
	case ChoseongPhieuph:
		return 'ㅍ'
	case ChoseongHieuh:
		return 'ㅎ'
	}

	return 0
}

// Rune returns the rune for Unicode `Hangul Jamo` or `Hangul Compatibility Jamo`.
// If `compatibility` is true, it returns the rune for `Hangul Compatibility Jamo`.
// Otherwise, it returns the rune for `Hangul Jamo`.
func (c Choseong) Rune(compatibility bool) rune {
	if compatibility {
		return c.CompatibilityRune()
	}
	return rune(c + ChoseongBase)
}

// Runes returns the slice of runes for Unicode `Hangul Jamo` or `Hangul Compatibility Jamo`.
// If `compatibility` is true, it returns the slice of runes for `Hangul Compatibility Jamo`.
// Otherwise, it returns the slice of runes for `Hangul Jamo`.
// If `level` is FullDecompose, it decomposes a double consonants into two to three runes.
func (c Choseong) Runes(compatibility bool, level int) []rune {
	if level < DefaultDecompose || level > FullDecompose {
		return nil
	}

	if level == DefaultDecompose {
		return []rune{c.Rune(compatibility)}
	}

	// FullDecompose
	switch c {
	case ChoseongSsangKiyeok:
		return []rune{ChoseongKiyeok.Rune(compatibility), ChoseongKiyeok.Rune(compatibility)}
	case ChoseongSsangTikeut:
		return []rune{ChoseongTikeut.Rune(compatibility), ChoseongTikeut.Rune(compatibility)}
	case ChoseongSsangPieup:
		return []rune{ChoseongPieup.Rune(compatibility), ChoseongPieup.Rune(compatibility)}
	case ChoseongSsangSios:
		return []rune{ChoseongSios.Rune(compatibility), ChoseongSios.Rune(compatibility)}
	case ChoseongSsangCieuc:
		return []rune{ChoseongCieuc.Rune(compatibility), ChoseongCieuc.Rune(compatibility)}
	}

	// Nothing to decompose
	return []rune{c.Rune(compatibility)}
}

// Jungseong is a medial vowel in Hangul.
// Unicode `Hangul Jamo` range for Jungseong: 0x1161 ~ 0x1175
// Unicode `Hangul Compatibility Jamo` range for Jungseong: 0x314F ~ 0x3163
// ref: https://en.wikipedia.org/wiki/Korean_language_and_computers#Hangul_in_Unicode
type Jungseong int32

const (
	JungseongA   = Jungseong(iota) // ㅏ
	JungseongAe                    // ㅐ
	JungseongYa                    // ㅑ
	JungseongYae                   // ㅒ
	JungseongEo                    // ㅓ
	JungseongE                     // ㅔ
	JungseongYeo                   // ㅕ
	JungseongYe                    // ㅖ
	JungseongO                     // ㅗ
	JungseongWa                    // ㅘ
	JungseongWae                   // ㅙ
	JungseongOe                    // ㅚ
	JungseongYo                    // ㅛ
	JungseongU                     // ㅜ
	JungseongWeo                   // ㅝ
	JungseongWe                    // ㅞ
	JungseongWi                    // ㅟ
	JungseongYu                    // ㅠ
	JungseongEu                    // ㅡ
	JungseongYi                    // ㅢ
	JungseongI                     // ㅣ

	JungseongBase  = 0x1161
	JungseongCount = 21
)

// CompatibilityRune returns the rune for Unicode `Hangul Compatibility Jamo`.
func (j Jungseong) CompatibilityRune() rune {
	switch j {
	case JungseongA:
		return 'ㅏ'
	case JungseongAe:
		return 'ㅐ'
	case JungseongYa:
		return 'ㅑ'
	case JungseongYae:
		return 'ㅒ'
	case JungseongEo:
		return 'ㅓ'
	case JungseongE:
		return 'ㅔ'
	case JungseongYeo:
		return 'ㅕ'
	case JungseongYe:
		return 'ㅖ'
	case JungseongO:
		return 'ㅗ'
	case JungseongWa:
		return 'ㅘ'
	case JungseongWae:
		return 'ㅙ'
	case JungseongOe:
		return 'ㅚ'
	case JungseongYo:
		return 'ㅛ'
	case JungseongU:
		return 'ㅜ'
	case JungseongWeo:
		return 'ㅝ'
	case JungseongWe:
		return 'ㅞ'
	case JungseongWi:
		return 'ㅟ'
	case JungseongYu:
		return 'ㅠ'
	case JungseongEu:
		return 'ㅡ'
	case JungseongYi:
		return 'ㅢ'
	case JungseongI:
		return 'ㅣ'
	}

	return 0
}

// Rune returns the rune for Unicode `Hangul Jamo` or `Hangul Compatibility Jamo`.
// If `compatibility` is true, it returns the rune for `Hangul Compatibility Jamo`.
// Otherwise, it returns the rune for `Hangul Jamo`.
func (j Jungseong) Rune(compatibility bool) rune {
	if compatibility {
		return j.CompatibilityRune()
	}
	return rune(j + JungseongBase)
}

// Runes returns the slice of runes for Unicode `Hangul Jamo` or `Hangul Compatibility Jamo`.
// If `compatibility` is true, it returns the slice of runes for `Hangul Compatibility Jamo`.
// Otherwise, it returns the slice of runes for `Hangul Jamo`.
// If `level` is FullDecompose, it decomposes a double vowels into two to three runes.
func (j Jungseong) Runes(compatibility bool, level int) []rune {
	if level < DefaultDecompose || level > FullDecompose {
		return nil
	}

	if level == DefaultDecompose {
		return []rune{j.Rune(compatibility)}
	}

	// FullDecompose
	switch j {
	case JungseongAe:
		return []rune{JungseongA.Rune(compatibility), JungseongI.Rune(compatibility)}
	case JungseongYae:
		return []rune{JungseongYa.Rune(compatibility), JungseongI.Rune(compatibility)}
	case JungseongE:
		return []rune{JungseongEo.Rune(compatibility), JungseongI.Rune(compatibility)}
	case JungseongYe:
		return []rune{JungseongYeo.Rune(compatibility), JungseongI.Rune(compatibility)}
	case JungseongWa:
		return []rune{JungseongO.Rune(compatibility), JungseongA.Rune(compatibility)}
	case JungseongWae:
		return []rune{JungseongO.Rune(compatibility), JungseongA.Rune(compatibility), JungseongI.Rune(compatibility)}
	case JungseongOe:
		return []rune{JungseongO.Rune(compatibility), JungseongI.Rune(compatibility)}
	case JungseongWeo:
		return []rune{JungseongU.Rune(compatibility), JungseongEo.Rune(compatibility)}
	case JungseongWe:
		return []rune{JungseongU.Rune(compatibility), JungseongEo.Rune(compatibility), JungseongI.Rune(compatibility)}
	case JungseongWi:
		return []rune{JungseongU.Rune(compatibility), JungseongI.Rune(compatibility)}
	case JungseongYi:
		return []rune{JungseongEu.Rune(compatibility), JungseongI.Rune(compatibility)}
	}

	// Nothing to decompose
	return []rune{j.Rune(compatibility)}
}

// Jongseong is a final consonant in Hangul.
// Unicode `Hangul Jamo` range for Jongseong: U+11A8 ~ U+11C2
// Unicode `Hangul Compatibility Jamo` range for Jongseong: U+3131 ~ U+314E
// ref: https://en.wikipedia.org/wiki/Korean_language_and_computers#Hangul_in_Unicode
type Jongseong int32 // 0x11A8 ~ 0x11C2

const (
	JongseongNull         = Jongseong(iota) // Blank
	JongseongKiyeok                         // ㄱ
	JongseongSsangKiyeok                    // ㄲ
	JongseongKiyeokSios                     // ㄳ
	JongseongNieun                          // ㄴ
	JongseongNieunCieuc                     // ㄵ
	JongseongNieunHieuh                     // ㄶ
	JongseongTikeut                         // ㄷ
	JongseongRieul                          // ㄹ
	JongseongRieulKiyeok                    // ㄺ
	JongseongRieulMieum                     // ㄻ
	JongseongRieulPieup                     // ㄼ
	JongseongRieulSios                      // ㄽ
	JongseongRieulThieuth                   // ㄾ
	JongseongRieulHieuh                     // ㅀ
	JongseongRieulPhieuph                   // ㄿ
	JongseongMieum                          // ㅁ
	JongseongPieup                          // ㅂ
	JongseongPieupSios                      // ㅄ
	JongseongSios                           // ㅅ
	JongseongSsangSios                      // ㅆ
	JongseongIeung                          // ㅇ
	JongseongCieuc                          // ㅈ
	JongseongChieuch                        // ㅊ
	JongseongKhieukh                        // ㅋ
	JongseongThieuth                        // ㅌ
	JongseongPhieuph                        // ㅍ
	JongseongHieuh                          // ㅎ

	JongseongBase  = 0x11A7
	JongseongCount = 28
)

// CompatibilityRune returns the rune for Unicode `Hangul Compatibility Jamo`.
func (j Jongseong) CompatibilityRune() rune {
	switch j {
	case JongseongKiyeok:
		return 'ㄱ'
	case JongseongSsangKiyeok:
		return 'ㄲ'
	case JongseongKiyeokSios:
		return 'ㄳ'
	case JongseongNieun:
		return 'ㄴ'
	case JongseongNieunCieuc:
		return 'ㄵ'
	case JongseongNieunHieuh:
		return 'ㄶ'
	case JongseongTikeut:
		return 'ㄷ'
	case JongseongRieul:
		return 'ㄹ'
	case JongseongRieulKiyeok:
		return 'ㄺ'
	case JongseongRieulMieum:
		return 'ㄻ'
	case JongseongRieulPieup:
		return 'ㄼ'
	case JongseongRieulSios:
		return 'ㄽ'
	case JongseongRieulThieuth:
		return 'ㄾ'
	case JongseongRieulPhieuph:
		return 'ㄿ'
	case JongseongRieulHieuh:
		return 'ㅀ'
	case JongseongMieum:
		return 'ㅁ'
	case JongseongPieup:
		return 'ㅂ'
	case JongseongPieupSios:
		return 'ㅄ'
	case JongseongSios:
		return 'ㅅ'
	case JongseongSsangSios:
		return 'ㅆ'
	case JongseongIeung:
		return 'ㅇ'
	case JongseongCieuc:
		return 'ㅈ'
	case JongseongChieuch:
		return 'ㅊ'
	case JongseongKhieukh:
		return 'ㅋ'
	case JongseongThieuth:
		return 'ㅌ'
	case JongseongPhieuph:
		return 'ㅍ'
	case JongseongHieuh:
		return 'ㅎ'
	}

	return 0
}

// Rune returns the rune for Unicode `Hangul Jamo` or `Hangul Compatibility Jamo`.
// If `compatibility` is true, it returns the rune for `Hangul Compatibility Jamo`.
// Otherwise, it returns the rune for `Hangul Jamo`.
func (j Jongseong) Rune(compatibility bool) rune {
	if j == JongseongNull {
		return 0
	}
	if compatibility {
		return j.CompatibilityRune()
	}
	return rune(j + JongseongBase)
}

// Runes returns the slice of runes for Unicode `Hangul Jamo` or `Hangul Compatibility Jamo`.
// If `compatibility` is true, it returns the slice of runes for `Hangul Compatibility Jamo`.
// Otherwise, it returns the slice of runes for `Hangul Jamo`.
// If `level` is FullDecompose, it decomposes a double consonants into two to three runes.
func (j Jongseong) Runes(compatibility bool, level int) []rune {
	if level < DefaultDecompose || level > FullDecompose {
		return nil
	}

	if j == JongseongNull {
		return nil
	}

	if level == DefaultDecompose {
		return []rune{j.Rune(compatibility)}
	}

	// FullDecompose
	switch j {
	case JongseongSsangKiyeok:
		return []rune{JongseongKiyeok.Rune(compatibility), JongseongKiyeok.Rune(compatibility)}
	case JongseongKiyeokSios:
		return []rune{JongseongKiyeok.Rune(compatibility), JongseongSios.Rune(compatibility)}
	case JongseongNieunCieuc:
		return []rune{JongseongNieun.Rune(compatibility), JongseongCieuc.Rune(compatibility)}
	case JongseongNieunHieuh:
		return []rune{JongseongNieun.Rune(compatibility), JongseongHieuh.Rune(compatibility)}
	case JongseongRieulKiyeok:
		return []rune{JongseongRieul.Rune(compatibility), JongseongKiyeok.Rune(compatibility)}
	case JongseongRieulMieum:
		return []rune{JongseongRieul.Rune(compatibility), JongseongMieum.Rune(compatibility)}
	case JongseongRieulPieup:
		return []rune{JongseongRieul.Rune(compatibility), JongseongPieup.Rune(compatibility)}
	case JongseongRieulSios:
		return []rune{JongseongRieul.Rune(compatibility), JongseongSios.Rune(compatibility)}
	case JongseongRieulThieuth:
		return []rune{JongseongRieul.Rune(compatibility), JongseongThieuth.Rune(compatibility)}
	case JongseongRieulPhieuph:
		return []rune{JongseongRieul.Rune(compatibility), JongseongPhieuph.Rune(compatibility)}
	case JongseongRieulHieuh:
		return []rune{JongseongRieul.Rune(compatibility), JongseongHieuh.Rune(compatibility)}
	case JongseongPieupSios:
		return []rune{JongseongPieup.Rune(compatibility), JongseongSios.Rune(compatibility)}
	case JongseongSsangSios:
		return []rune{JongseongSios.Rune(compatibility), JongseongSios.Rune(compatibility)}
	}

	// Nothing to decompose
	return []rune{j.Rune(compatibility)}
}
