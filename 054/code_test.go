package leetcode

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	var tests = []struct {
		matrix [][]int
		res    []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for _, tt := range tests {
		if res := spiralOrder(tt.matrix); !reflect.DeepEqual(res, tt.res) {
			t.Errorf("spiralOrder(%v) return %v, want %v", tt.matrix, res, tt.res)
		}
	}
}
