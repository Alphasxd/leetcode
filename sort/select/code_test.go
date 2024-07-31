package leetcode

import (
    "reflect"
    "testing"
)

func TestSelectSort(t *testing.T) {
    var tests = []struct {
        nums []int
        want []int
    }{
        {[]int{3, 2, 1}, []int{1, 2, 3}},
        {[]int{3, 2, 1, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
        {[]int{3, 2, 1, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
    }

    for _, tt := range tests {
        nums := make([]int, len(tt.nums))
        copy(nums, tt.nums)
        selectSort(tt.nums)
        if reflect.DeepEqual(tt.nums, tt.want) == false {
            t.Errorf("selectSort(%v) return %v, want %v", nums, tt.nums, tt.want)
        }

    }
}
