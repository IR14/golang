package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for m := map[float64]bool{}; !m[z]; z -= (z*z - x) / (2 * z) {
		m[z] = true
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(
		Sqrt(2),
		math.Sqrt(2),
	)
}

// 1
// 1.5
// 1.4166666666666667
// 1.4142156862745099
// 1.4142135623746899
// 1.4142135623730951
// 1.414213562373095
// 1.4142135623730951 1.4142135623730951
