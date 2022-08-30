package helper

func ArrCount[T int | string](arr []T) map[T]int {
	mapCount := make(map[T]int, 0)
	for _, val := range arr {
		mapCount[val]++
	}

	return mapCount
}
