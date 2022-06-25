package humanize

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var kb int = 1024
var mb int = kb * 1024
var gb int = mb * 1024
var tb int = gb * 1024

type kind struct {
	size   int
	symbol string
}

var denominations = [...]kind{{tb, "T"}, {gb, "G"}, {mb, "M"}, {kb, "K"}}

var linePattern = regexp.MustCompile(`\s+\d+$|^\d+\s+|\s+\d+\s+`)
var numberPattern = regexp.MustCompile(`\d+`)

func HumanizeNumber(s string) string {
	return linePattern.ReplaceAllStringFunc(s, replaceNumber)
}

func replaceNumber(s string) string {
	return numberPattern.ReplaceAllStringFunc(s, humanizeNumber)
}

func humanizeNumber(s string) string {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}

	for _, t := range denominations {
		if n >= t.size {
			f := float64(n) / float64(t.size)
			return fmt.Sprintf("%.2f%s", f, t.symbol)
		}
	}
	return strconv.Itoa(n)
}
