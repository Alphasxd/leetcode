package leetcode

import (
    "reflect"
    "testing"
)

func TestProductExceptSelf(t *testing.T) {
    var tests = []struct {
        nums     []int
        products []int
    }{
        {[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
        {[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
    }

    for _, tt := range tests {
        if products := productExceptSelf(tt.nums); !reflect.DeepEqual(products, tt.products) {
            t.Errorf("productExceptSelf(%v) return %v, want %v", tt.nums, products, tt.products)
        }
    }
}
