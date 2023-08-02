package leetcode

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	var tests = []struct {
		s, p string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
	}

	for _, tt := range tests {
		if want := findAnagrams(tt.s, tt.p); !reflect.DeepEqual(want, tt.want) {
			t.Errorf("findAnagrams(%v, %v) return %v, want %v", tt.s, tt.p, want, tt.want)
		}
	}
}