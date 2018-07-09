package luhn

import (
	"strconv"
)

func Valid(str string) bool {
	var arr []int
	for _, v := range str {
		if v == ' ' {
			continue
		}
		number, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
		arr = append(arr, number)
	}

	if len(arr) <= 1 {
		return false
	}
	k := len(arr) % 2 + 1
	sum := 0

	for _, number := range arr {
		r := (k % 2 + 1) * number
		sum += (r - 1) % 9 + 1
		k++
	}

	return sum % 10 == 0
}
