package maps

import "github.com/temaniarpit27/polygon-core-developer-course/helper"

func FindIntersection(arr1, arr2 []int) []int {
	if len(arr1) == 0 || len(arr2) == 0 {
		return make([]int, 0)
	}

	mapArr1 := helper.ArrCount(arr1)
	mapArr2 := helper.ArrCount(arr2)

	lenMapArr1 := len(mapArr1)
	lenMapArr2 := len(mapArr2)

	if lenMapArr1 < lenMapArr2 {
		return findIntersectionHelper(mapArr1, mapArr2)
	}

	return findIntersectionHelper(mapArr2, mapArr1)
}

func findIntersectionHelper(map1, map2 map[int]int) []int {
	resp := make([]int, 0)

	for k := range map1 {
		if _, ok := map2[k]; ok {
			resp = append(resp, k)
		}
	}

	return resp
}
