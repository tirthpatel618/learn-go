package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	change := float64(10)
	for math.Abs(change) > 0.01 {
		change = (z*z - x) / (2*z)
		z -= change
		fmt.Println(z)
	}
	return z
	
}

func main() {
	fmt.Println(Sqrt(100))
}
