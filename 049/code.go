// 49. 字母异位词分组
// https://leetcode-cn.com/problems/group-anagrams/
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		// 将字符串转换为字节切片，因为字符串是不可变的，而字节切片是可变的
		// 字符串是一个只读的 byte 类型切片
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}

	// 将 map 中的值转换为二维切片，因为 map 是无序的，所以需要遍历 map
	// 使用 make 创建一个二维切片，指定长度和容量，长度为 0，容量为 map 的长度 
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}

	sort.Slice(ans, func(i, j int) bool {
		return len(ans[i]) < len(ans[j])
	})
	
	return ans
}