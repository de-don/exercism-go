package scrabble

import (
	"strings"
)

func Score(str string) int {
	points := map[int]string{
		1:  "AEIOULNRST",
		2:  "DG",
		3:  "BCMP",
		4:  "FHVWY",
		5:  "K",
		8:  "JX",
		10: "QZ",
	}
	mapPoints := map[int32]int{}
	for k, v := range points {
		for _, letter := range v {
			mapPoints[letter] = k
		}
	}

	s := 0
	for _, i := range strings.ToUpper(str) {
		s += mapPoints[i]
	}
	return s
}
