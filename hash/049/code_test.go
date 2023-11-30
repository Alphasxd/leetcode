package leetcode

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	var tests = []struct {
		strs     []string
		anagrams [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"tan", "nat"}, {"eat", "tea", "ate"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}

	// 表格驱动测试
	for _, tt := range tests {
		anagrams := groupAnagrams(tt.strs)
		if reflect.DeepEqual(anagrams, tt.anagrams) == false {
			t.Errorf("groupAnagrams(%v) return %v, want %v", tt.strs, anagrams, tt.anagrams)
		}
	}
}
