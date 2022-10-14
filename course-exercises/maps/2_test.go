package maps

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumCount(t *testing.T) {
	t.Parallel()

	type test struct {
		inp []int
		op  map[int]int
	}

	tests := []test{
		{
			[]int{1, 2, 3, 4},
			map[int]int{
				1: 1,
				2: 1,
				3: 1,
				4: 1,
			},
		},
	}

	for _, tc := range tests {
		testOp := NumCount(tc.inp)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}
