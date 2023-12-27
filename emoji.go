package emoji

import (
	"unicode"

	"github.com/fiatjaf/emoji/data"
)

// ZWJ is the Unicode zero-width join character
const ZWJ = '\u200d'

// IsEmoji returns true if the specified rune has the (single-character)
// Emoji property in the latest Emoji version, false otherwise
func IsEmoji(r rune) bool {
	return unicode.Is(Emoji, r)
}

func isRegionalIndicator(r rune) bool {
	return unicode.Is(data.RegionalIndicator, r)
}

func isSkinToneModifier(r rune) bool {
	return unicode.Is(data.EmojiSkinToneModifier, r)
}

func isZeroWidth(r rune) bool {
	return r == ZWJ ||
		unicode.Is(unicode.Variation_Selector, r) ||
		unicode.Is(data.CombiningDiacritical, r) ||
		unicode.Is(data.Tag, r)
}

var Emoji = data.ParseRangeTable(data.BasicEmoji)
