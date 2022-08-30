package maps

import (
	"fmt"
	"sort"
)

type Account [20]byte

func SortData(inp map[int]map[int]Account) {
	keys1 := make([]int, 0)
	for k1 := range inp {
		keys1 = append(keys1, k1)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys1)))

	for _, k1 := range keys1 {
		keys2 := make([]int, 0)
		for k2 := range inp[k1] {
			keys2 = append(keys2, k2)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(keys2)))

		for _, k2 := range keys2 {
			fmt.Printf("Fee %d Nonce %d Account %s", k1, k2, inp[k1][k2])
		}

		fmt.Println()
	}
}
