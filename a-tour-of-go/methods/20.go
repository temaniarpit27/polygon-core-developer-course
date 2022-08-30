package methods

import (
	"fmt"
	"math"
)

type NegativeSqrtError float64

func (e NegativeSqrtError) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, NegativeSqrtError(x)
	}

	z := float64(1.)
	resp := float64(0)

	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(resp-z) < 1e-12 {
			break
		}

		resp = z
	}

	return resp, nil
}

func Methods20() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
