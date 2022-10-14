package maps

import (
	"strings"

	"github.com/temaniarpit27/polygon-core-developer-course/helper"
)

func WordCount(s string) map[string]int {
	if s == "" {
		return make(map[string]int, 0)
	}

	ss := strings.Fields(s)

	return helper.ArrCount(ss)
}
