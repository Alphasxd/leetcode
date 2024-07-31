package leetcode

import (
    "reflect"
    "testing"
)

func TestSetZeroes(t *testing.T) {
    var tests = []struct {
        matrix [][]int
        want   [][]int
    }{
        {[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
        {[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
    }

    for _, tt := range tests {
        before := make([][]int, len(tt.matrix))
        copy(before, tt.matrix)
        if setZeroes(tt.matrix); !reflect.DeepEqual(tt.matrix, tt.want) {
            t.Errorf("setZeroes(%v) return %v, want %v", before, tt.matrix, tt.want)
        }
    }
}
