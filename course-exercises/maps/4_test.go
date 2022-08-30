package maps

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonnaci(t *testing.T) {
	t.Parallel()

	type test struct {
		inp int
		op  int
	}

	tests := []test{
		{
			1,
			1,
		},
		{
			2,
			2,
		},
	}

	for _, tc := range tests {
		testOp := Fibonnaci(tc.inp)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}
