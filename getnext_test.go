package emoji

import (
	"testing"
)

func TestGetNext(t *testing.T) {
	for i, v := range []struct {
		text       string
		emojiChars int
	}{
		{"bla bla", 0},
		{"🇺🇸🇺🇸", 2},
		{"🇺🇸kjsd", 2},
		{"😋 zzz", 1},
		{"zzz 😋", 0},
		{"🏴󠁧󠁢󠁷󠁬󠁳󠁿 blob", 6},
		{"🏴󠁧󠁢󠁷󠁬󠁳󠁿", 6},
		{"😶‍🌫️", 4},
		{"🤽🏿‍♀️", 5},
		{"🦶🏼", 2},
		{"🦶🏼🦶🏼", 2},
		{" 🦶🏼🦶🏼", 0},
		{"👩🏼‍❤️‍👨🏼 askjdb", 8},
		{"👩🏼‍❤️‍👨🏼👩🏼‍❤️‍👨🏼 askjdb", 8},
		{"👩🏼‍❤️‍👨🏼👩🏼‍❤️‍👨🏼", 8},
		{"👩🏾‍🦳a", 4},
		{"🤙🏽☕️☀️;", 2},
		{"☕️☀️;", 2},
		{"☀️;", 2},
	} {
		runes := []rune(v.text)
		if count := GetNextEmojiCharacters(runes, len(runes)); count != v.emojiChars {
			t.Fatalf("%d: failed '%s' (%X) is %d, should be %d", i, v.text, runes, count, v.emojiChars)
		} else if len(GetNextEmoji(runes)) != count {
			t.Fatalf("%d: GetNextEmoji('%s' (%X)) is returning a different rune count", i, v.text, runes)
		}
	}
}
