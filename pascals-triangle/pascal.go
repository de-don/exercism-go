package pascal

func Triangle(n int) [][]int {
	result := [][]int{{1}}

	for i := 1; i < n; i++ {
		tmp := []int{1}
		for j := 1; j < i; j++ {
			tmp = append(tmp, result[i-1][j-1]+result[i-1][j])
		}
		tmp = append(tmp, 1)
		result = append(result, tmp)
	}
	return result
}
