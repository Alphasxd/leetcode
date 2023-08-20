package leetcode

import "testing"

func TestMaxSubArray(t *testing.T) {
	var tests = []struct {
		nums []int
		max  int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}

	for _, tt := range tests {
		if maxValue := maxSubArray(tt.nums); maxValue != tt.max {
			t.Errorf("maxSubArray(%v) return %v, want %v", tt.nums, maxValue, tt.max)
		}
	}
}
