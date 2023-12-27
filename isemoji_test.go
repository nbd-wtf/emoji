package emoji

import (
	"fmt"
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
		fmt.Printf("%d: '%s' (%X) \n", i, c, []rune(c))
		if !IsEmoji([]rune(c)[0]) {
			t.Fatalf("%d is an emoji, but got false", i)
		}
	}
}

// func TestIsNotEmoji(t *testing.T) {
// 	for i, c := range []string{
// 		"7",
// 		"8",
// 		"#",
// 		"*",
// 		"$",
// 		"¬",
// 	} {
// 		if IsEmoji([]rune(c)[0]) {
// 			t.Fatalf("%d: '%s' is not an emoji, but IsEmoji returned true", i, c)
// 		}
// 	}
// }
