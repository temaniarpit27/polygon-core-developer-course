package flowcontrol

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1.)
	resp := float64(0)

	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(resp-z) < 1e-12 {
			break
		}

		resp = z
	}

	return resp
}

func FlowControl8() {
	fmt.Println(Sqrt(2))
}
