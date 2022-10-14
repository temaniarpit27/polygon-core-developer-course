package maps

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIntersection(t *testing.T) {
	t.Parallel()

	type test struct {
		inp1 []int
		inp2 []int
		op   []int
	}

	tests := []test{
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2},
			[]int{1, 2},
		},
	}

	for _, tc := range tests {
		testOp := FindIntersection(tc.inp1, tc.inp2)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}
