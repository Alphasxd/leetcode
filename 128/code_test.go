package leetcode

import "testing"

func TestLongestConsecutive(t *testing.T) {
	var tests = []struct {
		nums   []int
		length int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}
	
	for _, tt := range tests {
		length := longestConsecutive(tt.nums)
		if length != tt.length {
			t.Errorf("longestConsecutive(%v) return %v, want %v", tt.nums, length, tt.length)
		}
	}
}