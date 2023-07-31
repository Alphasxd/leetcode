package leetcode

import "testing"

func TestTrap(t *testing.T) {
	var tests = []struct {
		height []int
		want   int
	}{
		{[]int{0,1,0,2,1,0,1,3,2,1,2,1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}

	for _, tt := range tests {
		want := trap(tt.height)
		if want != tt.want {
			t.Errorf("trap(%v) return error, want %d", tt.height, tt.want)
		}
	}
}
