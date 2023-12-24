// 242. 有效的字母异位词 https://leetcode-cn.com/problems/valid-anagram/

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

// 示例 1:
// 输入: s = "anagram", t = "nagaram"
// 输出: true

// 示例 2:
// 输入: s = "rat", t = "car"
// 输出: false

package leetcode

// import "sort"

// 只需要 s 和 t 中每个字符出现的次数都相同即可，和383题赎金信类似
// func isAnagram(s string, t string) bool {
// 	// 如果长度不相等，直接返回false
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	cnt := [26]int{}
// 	// 统计s中每个字符出现的次数
// 	for _, ch := range s {
// 		cnt[ch-'a']++
// 	}
// 	// 遍历t，减去s中出现的字符
// 	for _, ch := range t {
// 		cnt[ch-'a']--
// 		// 如果出现次数小于0，说明t中的字符在s中没有匹配的
// 		if cnt[ch-'a'] < 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// 排序，同个字母异位词排序后都是相同的
// func isAnagram(s, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	s1, s2 := []byte(s), []byte(t)
// 	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
// 	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
// 	return string(s1) == string(s2)
// }

// 考虑 rune，虽然题目要求是小写字母，但是可以扩展到unicode字符
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := make(map[rune]int)
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}