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
		"🔛",
		"👩",
		"🦳",
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

func TestCombinations(t *testing.T) {
	blackWhiteHairedWoman := []rune("👩🏾‍🦳")
	if len(blackWhiteHairedWoman) != 4 ||
		fmt.Sprintf("%X", blackWhiteHairedWoman[0]) != "1F469" ||
		fmt.Sprintf("%X", blackWhiteHairedWoman[1]) != "1F3FE" ||
		fmt.Sprintf("%X", blackWhiteHairedWoman[2]) != "200D" ||
		fmt.Sprintf("%X", blackWhiteHairedWoman[3]) != "1F9B3" ||
		!IsEmoji(blackWhiteHairedWoman[0]) ||
		!IsSkinToneModifier(blackWhiteHairedWoman[1]) ||
		!IsZeroWidthJoiner(blackWhiteHairedWoman[2]) ||
		!IsEmoji(blackWhiteHairedWoman[3]) {
		t.Fatalf("something wrong with the black white-haired woman: %s %X %X",
			string(blackWhiteHairedWoman), blackWhiteHairedWoman, blackWhiteHairedWoman[3])
	}

	couple := []rune("👩‍❤️‍👨")
	if len(couple) != 6 ||
		fmt.Sprintf("%X", couple[0]) != "1F469" ||
		fmt.Sprintf("%X", couple[1]) != "200D" ||
		fmt.Sprintf("%X", couple[2]) != "2764" ||
		fmt.Sprintf("%X", couple[3]) != "FE0F" ||
		fmt.Sprintf("%X", couple[4]) != "200D" ||
		fmt.Sprintf("%X", couple[5]) != "1F468" ||
		!IsEmoji(couple[0]) ||
		!IsZeroWidthJoiner(couple[1]) ||
		!IsEmoji(couple[2]) ||
		!IsVariationSelector(couple[3]) ||
		!IsZeroWidthJoiner(couple[4]) ||
		!IsEmoji(couple[5]) {
		t.Fatalf("something is wrong with the couple: %s %X", string(couple), couple)
	}

	whiteSuperWoman := []rune("🦸🏻‍♀️")
	if len(whiteSuperWoman) != 5 ||
		fmt.Sprintf("%X", whiteSuperWoman[0]) != "1F9B8" ||
		fmt.Sprintf("%X", whiteSuperWoman[1]) != "1F3FB" ||
		fmt.Sprintf("%X", whiteSuperWoman[2]) != "200D" ||
		fmt.Sprintf("%X", whiteSuperWoman[3]) != "2640" ||
		fmt.Sprintf("%X", whiteSuperWoman[4]) != "FE0F" ||
		!IsEmoji(whiteSuperWoman[0]) ||
		!IsSkinToneModifier(whiteSuperWoman[1]) ||
		!IsZeroWidthJoiner(whiteSuperWoman[2]) ||
		!IsEmoji(whiteSuperWoman[3]) ||
		!IsVariationSelector(whiteSuperWoman[4]) {
		t.Fatalf("something is wrong with the white super-woman: %s %X", string(whiteSuperWoman), whiteSuperWoman)
	}

	usFlag := []rune("🇺🇸")
	koreaFlag := []rune("🇰🇷")
	walesFlag := []rune("🏴󠁧󠁷󠁬󠁳󠁿")

	if len(usFlag) != 2 ||
		fmt.Sprintf("%X", usFlag[0]) != "1F1FA" ||
		fmt.Sprintf("%X", usFlag[1]) != "1F1F8" ||
		!IsRegionalIndicator(usFlag[0]) ||
		!IsRegionalIndicator(usFlag[1]) {
		t.Fatalf("something is wrong with the US flag: %s %X", string(usFlag), usFlag)
	}

	if len(koreaFlag) != 2 ||
		fmt.Sprintf("%X", koreaFlag[0]) != "1F1F0" ||
		fmt.Sprintf("%X", koreaFlag[1]) != "1F1F7" ||
		!IsRegionalIndicator(koreaFlag[0]) ||
		!IsRegionalIndicator(koreaFlag[1]) {
		t.Fatalf("something is wrong with the Korea flag: %s %X", string(koreaFlag), koreaFlag)
	}

	if len(walesFlag) != 6 ||
		fmt.Sprintf("%X", walesFlag[0]) != "1F3F4" ||
		fmt.Sprintf("%X", walesFlag[1]) != "E0067" ||
		fmt.Sprintf("%X", walesFlag[2]) != "E0077" ||
		fmt.Sprintf("%X", walesFlag[3]) != "E006C" ||
		fmt.Sprintf("%X", walesFlag[4]) != "E0073" ||
		fmt.Sprintf("%X", walesFlag[5]) != "E007F" ||
		!IsEmoji(walesFlag[0]) ||
		!IsTag(walesFlag[1]) ||
		!IsTag(walesFlag[2]) ||
		!IsTag(walesFlag[3]) ||
		!IsTag(walesFlag[4]) ||
		!IsTag(walesFlag[5]) {
		t.Fatalf("something is wrong with the Wales flag: %s %X", string(walesFlag), walesFlag)
	}
}
