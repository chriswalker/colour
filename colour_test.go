package colour

import (
	"fmt"
	"testing"
)

func TestColours(t *testing.T) {
	colours := []struct {
		text   string
		colour Colour
		fn     func(string) string
	}{
		{"black", FgBlack, Black},
		{"red", FgRed, Red},
		{"green", FgGreen, Green},
		{"yellow", FgYellow, Yellow},
		{"blue", FgBlue, Blue},
		{"magenta", FgMagenta, Magenta},
		{"cyan", FgCyan, Cyan},
		{"white", FgWhite, White},
		{"brblack", FgBrBlack, BrightBlack},
		{"brred", FgBrRed, BrightRed},
		{"brgreen", FgBrGreen, BrightGreen},
		{"bryellow", FgBrYellow, BrightYellow},
		{"brblue", FgBrBlue, BrightBlue},
		{"brmagenta", FgBrMagenta, BrightMagenta},
		{"brcyan", FgBrCyan, BrightCyan},
		{"brwhite", FgBrWhite, BrightWhite},
	}

	for _, c := range colours {

		expected := fmt.Sprintf("\x1b[%dm%s\x1b[0m", c.colour, c.text)
		formatted := c.fn(c.text)
		if expected != formatted {
			t.Errorf("Wrong color code; expected '%s', got '%s'", expected, formatted)
		}
	}
}
