package triangle

import (
	"math"
)

type Kind int

const (
	NaT = -1 // not a triangle
	Equ = 2  // equilateral
	Iso = 1  // isosceles
	Sca = 0  // scalene
)

func KindFromSides(a, b, c float64) Kind {
	x := a * b * c
	if (x == 0) || math.IsNaN(x) || math.IsInf(x, 0) {
		return NaT
	}
	if (a+b < c) || (a+c < b) || (b+c < a) {
		return NaT
	}
	var k Kind
	if a == b {
		k++
	}
	if a == c || c == b {
		k++
	}

	return k
}
