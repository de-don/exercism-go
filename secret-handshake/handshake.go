package secret

import (
	"fmt"
)

func Handshake(num uint) (result []string) {
	actions := []string{"wink", "double blink", "close your eyes", "jump"}

	digest := fmt.Sprintf("%05b", num)
	for i := 0; i < 4; i++ {
		if digest[4-i] == '1' {
			result = append(result, actions[i])
		}
	}
	n := len(result)
	for i := 0; i < n/2 && digest[0] == '1'; i++ {
		result[i], result[n-i-1] = result[n-i-1], result[i]
	}
	return result
}
