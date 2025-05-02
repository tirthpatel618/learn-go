package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %d", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	change := float64(10)
	for math.Abs(change) > 0.01 {
		change = (z*z - x) / (2*z)
		z -= change
	}
	return z, nil
}





func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
