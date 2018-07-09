package diffsquares

import "math"

func SquareOfSums(n int) int {
	s := 0
	for i := 0; i <= n; i++ {
		s += i
	}
	return int(math.Pow(float64(s), 2.0))
}

func SumOfSquares(n int) int {
	s := 0.0
	for i := 0; i <= n; i++ {
		s += math.Pow(float64(i), 2.0)
	}
	return int(s)
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
