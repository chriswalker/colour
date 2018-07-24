package colour

import "fmt"

type Colour int

const escape = "\x1b"

const (
	FgBlack = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

const (
	FgBrBlack = iota + 90
	FgBrRed
	FgBrGreen
	FgBrYellow
	FgBrBlue
	FgBrMagenta
	FgBrCyan
	FgBrWhite
)

// Reset value
const Reset = 0

func Black(s string) string { return colour(FgBlack, s) }

func Red(s string) string { return colour(FgRed, s) }

func Green(s string) string { return colour(FgGreen, s) }

func Yellow(s string) string { return colour(FgYellow, s) }

func Blue(s string) string { return colour(FgBlue, s) }

func Magenta(s string) string { return colour(FgMagenta, s) }

func Cyan(s string) string { return colour(FgCyan, s) }

func White(s string) string { return colour(FgWhite, s) }

func BrightBlack(s string) string { return colour(FgBrBlack, s) }

func BrightRed(s string) string { return colour(FgBrRed, s) }

func BrightGreen(s string) string { return colour(FgBrGreen, s) }

func BrightYellow(s string) string { return colour(FgBrYellow, s) }

func BrightBlue(s string) string { return colour(FgBrBlue, s) }

func BrightMagenta(s string) string { return colour(FgBrMagenta, s) }

func BrightCyan(s string) string { return colour(FgBrCyan, s) }

func BrightWhite(s string) string { return colour(FgBrWhite, s) }

// colour wraps the supplied string with the supplied colour and reset codes
func colour(c Colour, s string) string {
	return fmt.Sprintf("%s[%dm%s%s", escape, c, s, reset())
}

// reset returns a reset escape sequence
func reset() string {
	return fmt.Sprintf("%s[%dm", escape, Reset)
}
