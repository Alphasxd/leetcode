// 242. 有效的字母异位词 https://leetcode.cn/problems/valid-anagram/

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

// 示例 1:
// 输入: s = "anagram", t = "nagaram"
// 输出: true

// 示例 2:
// 输入: s = "rat", t = "car"
// 输出: false

package leetcode

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
