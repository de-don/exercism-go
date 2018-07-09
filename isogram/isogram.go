package isogram

import "strings"

func IsIsogram(str string) bool {
	flags := map[string]bool{}

	for _, letter := range str{
		l := strings.ToLower(string(letter))
		if !("a" <= l && l <= "z") &&
		   !("A" <= l && l <= "Z"){
			continue
		}
		if flags[l]{
			return false
		}
		flags[l] = true
	}

	return true
}
