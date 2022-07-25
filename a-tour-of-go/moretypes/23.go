package moretypes

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ss := strings.Fields(s)
	len_ss := len(ss)
	resp := make(map[string]int)

	for i := 0; i < len_ss; i++ {
		(resp[ss[i]])++
	}

	return resp
}

func MoreTypes23() {
	wc.Test(WordCount)
}
