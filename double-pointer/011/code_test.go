package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	var tests = []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, tt := range tests {
		want := maxArea(tt.height)
		if want != tt.want {
			t.Errorf("maxArea(%v) return %v, want %v", tt.height, want, tt.want)
		}
	}
}
