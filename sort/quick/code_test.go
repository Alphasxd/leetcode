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
        {[]int{35, 24, 48, 21, 43, 88, 75, 91, 90}, 0, 8, []int{21, 24, 35, 43, 48, 75, 88, 90, 91}},
        {[]int{45, 23, 89, 77, 32, 56, 29}, 0, 6, []int{23, 29, 32, 45, 56, 77, 89}},
        {[]int{67, 33, 14, 82, 26, 51, 94, 73, 28, 57}, 0, 9, []int{14, 26, 28, 33, 51, 57, 67, 73, 82, 94}},
        {[]int{7, 6, 5, 4, 3, 2, 1}, 0, 6, []int{1, 2, 3, 4, 5, 6, 7}},
        {[]int{150, 200, 250, 300, 350}, 0, 4, []int{150, 200, 250, 300, 350}},
        {[]int{5}, 0, 0, []int{5}},
        {[]int{}, 0, -1, []int{}},
        {[]int{8, 6, 4, 2, 0}, 0, 4, []int{0, 2, 4, 6, 8}},
        {[]int{20, 40, 60, 80, 100, 120, 140, 160, 180}, 0, 8, []int{20, 40, 60, 80, 100, 120, 140, 160, 180}},
        {[]int{180, 160, 140, 120, 100, 80, 60, 40, 20}, 0, 8, []int{20, 40, 60, 80, 100, 120, 140, 160, 180}},
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
