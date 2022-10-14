package maps

import "github.com/temaniarpit27/polygon-core-developer-course/helper"

func NumCount(arr []int) map[int]int {
	if len(arr) == 0 {
		return make(map[int]int, 0)
	}

	return helper.ArrCount(arr)
}
