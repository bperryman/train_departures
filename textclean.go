package train_departures

import (
	"regexp"
	"strings"
	"unicode"
)

// Numerous routines for the clearning up of text

func mapWhiteSpace(ws rune) rune {
	switch {
	case unicode.IsSpace(ws):
		return ' '
	case ws == 0xFFFD:
		return ' '
	default:
		return ws
	}
}

func cleanText(txt string) string {
	excessSpace := regexp.MustCompile(`\s{2,}`)
	returnText := strings.Map(mapWhiteSpace, txt)
	returnText = strings.Trim(returnText, " \t\r\n")
	return excessSpace.ReplaceAllLiteralString(returnText, " ")
}
