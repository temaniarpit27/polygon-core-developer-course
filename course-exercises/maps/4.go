package maps

func Fibonnaci(n int) int {
	fibMemo := make(map[int]int, 0)
	fibMemo[0] = 1
	fibMemo[1] = 1

	return fibonnaciHelper(n, fibMemo)
}

func fibonnaciHelper(n int, fibMemo map[int]int) int {
	if val, ok := fibMemo[n]; ok {
		return val
	}

	fibMemo[n] = fibonnaciHelper(n-1, fibMemo) + fibonnaciHelper(n-2, fibMemo)

	return fibMemo[n]
}
