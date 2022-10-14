package slices

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExercise1(t *testing.T) {
	t.Parallel()

	type test struct {
		inp1 []int
		inp2 int
		op   []int
	}

	tests := []test{
		{
			[]int{1, 2},
			1,
			[]int{2, 3},
		},
	}

	for _, tc := range tests {
		testOp := AddElement(tc.inp1, tc.inp2)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}

func TestExercise2(t *testing.T) {
	t.Parallel()

	type test struct {
		inp1 []int
		inp2 int
		op   []int
	}

	tests := []test{
		{
			[]int{1, 2},
			1,
			[]int{1, 2, 1},
		},
	}

	for _, tc := range tests {
		testOp := AddElementToEnd(tc.inp1, tc.inp2)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}

func TestExercise3(t *testing.T) {
	t.Parallel()

	type test struct {
		inp1 []int
		inp2 int
		op   []int
	}

	tests := []test{
		{
			[]int{1, 2},
			1,
			[]int{1, 1, 2},
		},
	}

	for _, tc := range tests {
		testOp := AddElementToStart(tc.inp1, tc.inp2)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}

func TestExercise456(t *testing.T) {
	t.Parallel()

	type test struct {
		inp1 []int
		inp2 int
		op1  int
		op2  []int
	}

	tests := []test{
		// 4
		{
			[]int{1, 2, 3, 4},
			3,
			4,
			[]int{1, 2, 3},
		},
		// 5
		{
			[]int{1, 2, 3, 4},
			0,
			1,
			[]int{2, 3, 4},
		},
		// 6
		{
			[]int{1, 2, 3, 4},
			1,
			2,
			[]int{1, 3, 4},
		},
	}

	for _, tc := range tests {
		testOp1, testOp2 := RemoveElement(tc.inp1, tc.inp2)
		assert.True(t, reflect.DeepEqual(testOp1, tc.op1))
		assert.True(t, reflect.DeepEqual(testOp2, tc.op2))
	}
}

func TestExercise7(t *testing.T) {
	t.Parallel()

	type test struct {
		inp1 []int
		inp2 []int
		op   []int
	}

	tests := []test{
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3},
			[]int{4},
		},
	}

	for _, tc := range tests {
		testOp := RemoveDuplicates(tc.inp1, tc.inp2)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}

func TestExercise8(t *testing.T) {
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
			[]int{3, 4},
		},
	}

	for _, tc := range tests {
		testOp := RemoveDuplicatesFromSecond(tc.inp1, tc.inp2)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}

func TestExercise910(t *testing.T) {
	t.Parallel()

	type test struct {
		inp1 []int
		inp2 int
		op   []int
	}

	tests := []test{
		// 9
		{
			[]int{1, 2, 3, 4},
			1,
			[]int{2, 3, 4, 1},
		},
		// 10
		{
			[]int{1, 2, 3, 4},
			2,
			[]int{3, 4, 1, 2},
		},
	}

	for _, tc := range tests {
		testOp := RotateLeft(tc.inp1, tc.inp2)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}

func TestExercise1112(t *testing.T) {
	t.Parallel()

	type test struct {
		inp1 []int
		inp2 int
		op   []int
	}

	tests := []test{
		// 11
		{
			[]int{1, 2, 3, 4},
			1,
			[]int{4, 1, 2, 3},
		},
		// 12
		{
			[]int{1, 2, 3, 4},
			2,
			[]int{3, 4, 1, 2},
		},
	}

	for _, tc := range tests {
		testOp := RotateRight(tc.inp1, tc.inp2)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}

func TestExercise13(t *testing.T) {
	t.Parallel()

	type test struct {
		inp []int
		op  []int
	}

	tests := []test{
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		testOp := Copy(tc.inp)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}

func TestExercise14(t *testing.T) {
	t.Parallel()

	type test struct {
		inp []int
		op  []int
	}

	tests := []test{
		{
			[]int{1, 2, 3, 4},
			[]int{2, 1, 4, 3},
		},
		{
			[]int{1, 2, 3},
			[]int{2, 1, 3},
		},
	}

	for _, tc := range tests {
		testOp := SwapOddEvenElem(tc.inp)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}

func TestExercise15(t *testing.T) {
	t.Parallel()

	type test1 struct {
		inp1 []int
		inp2 int
		op   []int
	}

	tests := []test1{
		{
			[]int{4, 3, 2, 1},
			1,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 2, 3, 4},
			2,
			[]int{4, 3, 2, 1},
		},
		{
			[]int{1, 3, 11, 4},
			3,
			[]int{1, 11, 3, 4},
		},
	}

	for _, tc := range tests {
		testOp := SortSliceOrder(tc.inp1, tc.inp2)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}

	type test2 struct {
		inp1 []string
		inp2 int
		op   []string
	}

	tests2 := []test2{
		{
			[]string{"Hey", "Hey2"},
			1,
			[]string{"Hey", "Hey2"},
		},
		{
			[]string{"Hey", "Hey2"},
			2,
			[]string{"Hey2", "Hey"},
		},
		{
			[]string{"Hey2", "Hey"},
			3,
			[]string{"Hey", "Hey2"},
		},
	}

	for _, tc := range tests2 {
		testOp := SortSliceOrder(tc.inp1, tc.inp2)
		assert.True(t, reflect.DeepEqual(testOp, tc.op))
	}
}
