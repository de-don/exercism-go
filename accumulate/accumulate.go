package accumulate

func Accumulate(data []string, converter func(string) string) []string {
	result := make([]string, len(data))

	for i, v := range data {
		result[i] = converter(v)
	}
	return result
}
