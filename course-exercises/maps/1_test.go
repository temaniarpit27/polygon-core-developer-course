package maps

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	t.Parallel()

	type test struct {
		inp string
		op  map[string]int
	}

	tests := []test{
		{
			"Hey",
			map[string]int{
				"Hey": 1,
			},
		},
	}

	for _, tc := range tests {
		testOp := WordCount(tc.inp)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}
