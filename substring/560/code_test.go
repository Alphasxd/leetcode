package leetcode

import "testing"

func TestSubarrySum(t *testing.T) {
    var tests = []struct {
        nums []int
        k    int
        want int
    }{
        {[]int{1, 1, 1}, 2, 2},
        {[]int{1, 2, 3}, 3, 2},
    }

    for _, tt := range tests {
        if want := subarraySum(tt.nums, tt.k); want != tt.want {
            t.Errorf("subarraySum(%v, %v) return %v, want %v", tt.nums, tt.k, want, tt.want)
        }
    }
}
