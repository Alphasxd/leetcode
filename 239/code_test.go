package leetcode

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{}, 0, []int{}},
	}

	for _, tt := range tests {
		if want := maxSlidingWindow(tt.nums, tt.k); !reflect.DeepEqual(want, tt.want) {
			t.Errorf("maxSlidingWindow(%v, %v) return %v, want %v", tt.nums, tt.k, want, tt.want)
		}
	}
}
