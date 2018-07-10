package series

func All(n int, s string) (result []string) {
	for i := 0; i < len(s)-n+1; i++ {
		result = append(result, s[i:i+n])
	}
	return result
}

func UnsafeFirst(n int, s string) (first string) {
	return s[:n]
}

func First(n int, str string) (first string, ok bool) {
	if len(str) < n {
		return "", false
	}
	return UnsafeFirst(n, str), true
}
