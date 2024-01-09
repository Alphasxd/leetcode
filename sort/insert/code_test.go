package leetcode

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{5, 2, 4, 6, 1, 3},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			input: []int{5, 3, 8, 4, 2},
			want:  []int{2, 3, 4, 5, 8},
		},
		{
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		nums := make([]int, len(test.input))
		copy(nums, test.input)
		insertionSort(test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("insertionSort(%v) => %v, want %v", nums, test.input, test.want)
		}
	}
}