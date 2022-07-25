package moretypes

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var fib1, fib2 = 0, 1

	var fib3 = fib1 + fib2

	return func() int {
		fib3, fib1, fib2 = fib1, fib2, fib1+fib2
		return fib3
	}
}

func MoreTypes26() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
