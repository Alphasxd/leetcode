package leetcode

import (
    "reflect"
    "testing"
)

func TestThreeSum(t *testing.T) {
    var tests = []struct {
        nums     []int
        solution [][]int
    }{
        {[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
        {[]int{0, 1, 1}, [][]int{}},
        {[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
    }

    for _, tt := range tests {
        solution := threeSum(tt.nums)
        if !reflect.DeepEqual(solution, tt.solution) {
            t.Errorf("threeSum(%v) return %v, want %v", tt.nums, solution, tt.solution)
        }
    }
}
