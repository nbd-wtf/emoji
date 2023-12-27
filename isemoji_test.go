package emoji

import (
	"testing"
)

func TestIsEmoji(t *testing.T) {
	for i, c := range []string{
		"⛩️",
		"⌚",
		"☕",
		"☕️",
		"🏽",
		"☀️",
		"🍁",
	} {
		if !IsEmoji([]rune(c)[0]) {
			t.Fatalf("%d: '%s' (%X) is an emoji, but got false", i, c, []rune(c))
		}
	}
}

func TestIsNotEmoji(t *testing.T) {
	for i, c := range []string{
		"7",
		"8",
		"#",
		"*",
		"$",
		"¬",
	} {
		if IsEmoji([]rune(c)[0]) {
			t.Fatalf("%d: '%s' is not an emoji, but IsEmoji returned true", i, c)
		}
	}
}
