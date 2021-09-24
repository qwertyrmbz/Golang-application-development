package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	z_before := 0.0
	count := 0
	for i := 0; math.Abs(z-z_before) > 0.0000001; i++ {
		z_before = z
		z = z - (math.Pow(z, 2)-x)/(2*z)
		count += 1
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}