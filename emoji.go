package emoji

import (
	"unicode"

	"github.com/fiatjaf/emoji/data"
)

var emoji = data.ParseRangeTable(data.BasicEmoji)

// this is the range of base emojis, it also includes things that are used only in combinations,
// like "the red hair"
func IsEmoji(r rune) bool { return unicode.Is(emoji, r) }

// skin tone modifiers are applied right after emojis to change their skin tones
func IsSkinToneModifier(r rune) bool { return unicode.Is(data.EmojiSkinToneModifier, r) }

// the ZWJ is used to combine emojis into a single image, like "couple + man + woman"
func IsZeroWidthJoiner(r rune) bool { return r == '\u200d' }

// the variation selector is a useless thing that tells the renderer if it should display an
// image or a text format of the emoji, can be safely ignored
func IsVariationSelector(r rune) bool { return r == '\uFE0F' || r == '\uFE0E' }

// regional indicators are used to make flags: each regional indicator represents a letter,
// if you combine one that represents "B" with one that represents "R" you get the flag of
// Brazil
func IsRegionalIndicator(r rune) bool { return unicode.Is(data.RegionalIndicator, r) }

// tags are used for subnational flags like the flag of Wales, they consist of the black flag emoji
// followed by 5 tag runes that work like the country flags, but specifying the ISO-something
// code for subnational regions (Wales is "gbwls")
func IsTag(r rune) bool { return unicode.Is(data.Tag, r) }
