package leetcode

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	var tests = []struct {
		intervals [][]int
		merged    [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 3}, {2, 4}, {4, 8}}, [][]int{{1, 8}}},
	}

	for _, tt := range tests {
		before := make([][]int, len(tt.intervals))
		copy(before, tt.intervals)
		if merged := merge(tt.intervals); !reflect.DeepEqual(merged, tt.merged) {
			t.Errorf("merge(%v) return %v, want %v", before, merged, tt.merged)
		}
	}
}
