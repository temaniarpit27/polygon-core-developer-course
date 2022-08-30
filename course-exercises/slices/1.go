package slices

import (
	"fmt"
	"sort"

	"github.com/temaniarpit27/polygon-core-developer-course/helper"
)

func Exercise1(arr []int, n int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] += n
	}

	return arr
}

func Exercise2(arr []int, n int) []int {
	arr = append(arr, n)
	return arr
}

func Exercise3(arr []int, n int) []int {
	resp := make([]int, 0)
	resp = append(resp, n)
	resp = append(resp, arr...)

	return resp
}

func Exercise456(arr []int, index int) (int, []int) {
	elem := arr[index]
	return elem, append(arr[:index], arr[index+1:]...)
}

func Exercise7(arr1 []int, arr2 []int) []int {
	arrCount1 := helper.ArrCount(arr1)
	arrCount2 := helper.ArrCount(arr2)
	mergedArr := make([]int, 0)

	for k := range arrCount1 {
		if _, ok := arrCount2[k]; !ok {
			mergedArr = append(mergedArr, k)
		}
	}

	for k := range arrCount2 {
		if _, ok := arrCount1[k]; !ok {
			mergedArr = append(mergedArr, k)
		}
	}

	return mergedArr
}

func Exercise8(arr1 []int, arr2 []int) []int {
	arrCount2 := helper.ArrCount(arr2)
	for i := 0; i < len(arr1); i++ {
		if _, ok := arrCount2[arr1[i]]; ok {
			_, arr1 = Exercise456(arr1, i)
			i = i - 1
		}
	}

	return arr1
}

func Exercise910(arr []int, index int) []int {
	return append(arr[index:], arr[:index]...)
}

func Exercise1112(arr []int, index int) []int {
	return append(arr[len(arr)-index:], arr[:len(arr)-index]...)
}

func Exercise13(arr []int) []int {
	arr2 := make([]int, len(arr))
	copy(arr2, arr)

	return arr2
}

func Exercise14(arr []int) []int {
	for i := 0; i < len(arr); i += 2 {
		if i == len(arr)-1 {
			break
		}

		arr[i], arr[i+1] = arr[i+1], arr[i]
	}

	return arr
}

// method -> 1 - direct, 2 - reverse, 3 - lexicographic
func Exercise15[T int | string](arr []T, method int) []T {
	sort.Slice(arr, func(i, j int) bool {
		if method == 1 {
			return arr[i] < arr[j]
		} else if method == 2 {
			return arr[i] > arr[j]
		} else {
			return fmt.Sprint(arr[i]) < fmt.Sprint(arr[j])
		}
	})

	return arr
}
