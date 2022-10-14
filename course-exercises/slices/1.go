package slices

import (
	"fmt"
	"sort"

	"github.com/temaniarpit27/polygon-core-developer-course/helper"
)

func AddElement(arr []int, n int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] += n
	}

	return arr
}

func AddElementToEnd(arr []int, n int) []int {
	arr = append(arr, n)
	return arr
}

func AddElementToStart(arr []int, n int) []int {
	resp := make([]int, len(arr)+1)
	resp[0] = n
	copy(resp[1:], arr)

	return resp
}

func RemoveElement(arr []int, index int) (int, []int) {
	elem := arr[index]
	return elem, append(arr[:index], arr[index+1:]...)
}

func RemoveDuplicates(arr1 []int, arr2 []int) []int {
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

func RemoveDuplicatesFromSecond(arr1 []int, arr2 []int) []int {
	arrCount2 := helper.ArrCount(arr2)
	for i := 0; i < len(arr1); i++ {
		if _, ok := arrCount2[arr1[i]]; ok {
			_, arr1 = RemoveElement(arr1, i)
			i = i - 1
		}
	}

	return arr1
}

func RotateLeft(arr []int, index int) []int {
	return append(arr[index:], arr[:index]...)
}

func RotateRight(arr []int, index int) []int {
	return append(arr[len(arr)-index:], arr[:len(arr)-index]...)
}

func Copy(arr []int) []int {
	arr2 := make([]int, len(arr))
	copy(arr2, arr)

	return arr2
}

func SwapOddEvenElem(arr []int) []int {
	for i := 0; i < len(arr); i += 2 {
		if i == len(arr)-1 {
			break
		}

		arr[i], arr[i+1] = arr[i+1], arr[i]
	}

	return arr
}

// method -> 1 - direct, 2 - reverse, 3 - lexicographic
func SortSliceOrder[T int | string](arr []T, method int) []T {
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
