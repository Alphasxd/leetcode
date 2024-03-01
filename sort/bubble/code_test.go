package leetcode

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var tests = []struct {
		nums []int
		sorted []int
	} {
		{[]int{49, 38, 65, 97, 76, 13, 27, 49}, []int{13, 27, 38, 49, 49, 65, 76, 97}},
		{[]int{3, 2, 1, 5, 6, 4}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{28, 77, 98, 26, 33, 51, 20, 14, 88, 67}, []int{14, 20, 26, 28, 33, 51, 67, 77, 88, 98}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)
		bubbleSort(tt.nums)
		if reflect.DeepEqual(tt.nums, tt.sorted) == false {
			t.Errorf("bubbleSort(%v) return %v, want %v", nums, tt.nums, tt.sorted)
		}
	}
}