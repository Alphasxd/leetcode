package leetcode

import (
    "reflect"
    "testing"
)

func TestMerge(t *testing.T) {
    var tests = []struct {
        nums1 []int
        m     int
        nums2 []int
        n     int
        nums  []int
    }{
        {[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
        {[]int{1}, 1, []int{}, 0, []int{1}},
        {[]int{0}, 0, []int{1}, 1, []int{1}},
    }

    for _, tt := range tests {
        // 保留 nums1 的原始副本
        nums1 := make([]int, len(tt.nums1))
        copy(nums1, tt.nums1)

        merge(tt.nums1, tt.m, tt.nums2, tt.n)
        if reflect.DeepEqual(tt.nums1, tt.nums) == false {
            t.Errorf("merge(%v, %v, %v, %v) return %v, want %v", nums1, tt.m, tt.nums2, tt.n, tt.nums1, tt.nums)
        }
    }
}
