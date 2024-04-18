// 49. 字母异位词分组 https://leetcode-cn.com/problems/group-anagrams/

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

// 示例 1:
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

// 示例 2:
// 输入: strs = [""]
// 输出: [[""]]

package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		// 将字符串转换为字节切片，因为字符串是不可变的，而字节切片是可变的
		// 字符串是一个只读的 byte 类型切片
		s := []byte(str)
		// 对字节切片进行排序, 按字母顺序排序
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}

	// 将 map 中的值转换为二维切片，因为 map 是无序的，所以需要遍历 map
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}

	return ans
}
