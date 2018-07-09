package robotname

import (
	"strconv"
	"math/rand"
)

type Robot struct {
	name string
}

func genRandName() string {
	letter1 := string(rune(rand.Intn(26) + int('A')))
	letter2 := string(rune(rand.Intn(26) + int('A')))
	number := strconv.Itoa(rand.Intn(899) + 100)
	return letter1 + letter2 + number
}

var usedNames map[string]bool

func (r *Robot) Name() string {
	for ; len(r.name) == 0; {
		name := genRandName()
		if usedNames == nil {
			usedNames = make(map[string]bool)
		}
		if !usedNames[name] {
			usedNames[name] = true
			r.name = name
			break
		}
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = ""
}
