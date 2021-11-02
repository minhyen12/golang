package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	for z := 1.0; ; {
		state := z
		z -= (z*z - x) / (2 * z)
		if float32(state) == float32(z) {
			return z
		}
	}
}

func main() {
	p := float64(0.000001)
	fmt.Println(Sqrt(p), Sqrt(p) == math.Sqrt(p))
}
