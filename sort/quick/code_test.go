package leetcode

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var tests = []struct {
		nums   []int
		lo     int
		hi     int
		sorted []int
	}{
		{[]int{49, 38, 65, 97, 76, 13, 27, 49}, 0, 7, []int{13, 27, 38, 49, 49, 65, 76, 97}},
		{[]int{3, 2, 1, 5, 6, 4}, 0, 5, []int{1, 2, 3, 4, 5, 6}},
		{[]int{28, 77, 98, 26, 33, 51, 20, 14, 88, 67}, 0, 9, []int{14, 20, 26, 28, 33, 51, 67, 77, 88, 98}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)
		quickSort(tt.nums, tt.lo, tt.hi)
		if reflect.DeepEqual(tt.nums, tt.sorted) == false {
			t.Errorf("quickSort(%v, %v, %v) return %v, want %v", nums, tt.lo, tt.hi, tt.nums, tt.sorted)
		}
	}
}
