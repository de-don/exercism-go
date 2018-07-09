package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	re := regexp.MustCompile(`\W`)
	s = re.ReplaceAllString(s, " ")
	words := strings.Split(s, " ")

	s = ""

	for i := range words {
		if words[i] != "" {
			s += string(words[i][0])
		}
	}

	return strings.ToUpper(s)
}
