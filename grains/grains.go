package grains

import (
	"errors"
	"math"
	"math/big"
)

func Square(num int) (uint64, error) {
	if num <= 0 || num > 64{
		return 0, errors.New("invalid number")
	}

	return uint64(math.Pow(float64(2), float64(num - 1))), nil
}

func Total() uint64 {
	var i, e = big.NewInt(2), big.NewInt(64)
	i.Exp(i, e, nil)
	i.Sub(i, big.NewInt(1))
	return i.Uint64()
}