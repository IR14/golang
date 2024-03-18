package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	for m := map[float64]bool{}; !m[z]; z -= (z*z - x) / (2 * z) {
		m[z] = true
		fmt.Println(z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

// 1
// 1.5
// 1.4166666666666667
// 1.4142156862745099
// 1.4142135623746899
// 1.4142135623730951
// 1.414213562373095
// 1.4142135623730951 <nil>
// -2 cannot Sqrt negative number: -2
