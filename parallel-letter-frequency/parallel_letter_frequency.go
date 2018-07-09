package letter

import "time"

func ConcurrentFrequency(s []string) FreqMap {
	ch := make(chan rune)
	m := FreqMap{}

	for _, str := range s {
		go func(str string) {
			for _, r := range str {
				ch <- r
			}
		}(str)
	}
	for {
		select {
		case r := <-ch:
			m[r]++
		case <-time.After(1 * time.Second):
			return m
		}
	}
}
