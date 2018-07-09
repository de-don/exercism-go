package bob

import (
	"strings"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}

	yell := (strings.ToUpper(remark) == remark) && !(strings.ToUpper(remark) == strings.ToLower(remark))
	question := remark[len(remark) - 1] == '?'
	if yell && question {
		return "Calm down, I know what I'm doing!"
	}
	if !yell && question {
		return "Sure."
	}
	if yell && !question {
		return "Whoa, chill out!"
	}

	return "Whatever."
}