package acronym

const testVersion = 3

import (
	"strings"
	"unicode"
)

func abbreviate(s string) string {
	acronym := ""
	for i, l := range s {
		sl := strings.ToUpper(string(l))
		if i == 0 {
			acronym += sl
			continue
		}
		prev := rune(s[i-1])
		curr := rune(s[i])
		letterAfterNonLetter := unicode.IsLetter(curr) && !unicode.IsLetter(prev)
		upperAfterNonUpper := unicode.IsUpper(curr) && !unicode.IsUpper(prev)
		if letterAfterNonLetter || upperAfterNonUpper {
			acronym += sl
		}
	}
	return acronym
}
