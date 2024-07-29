// 5. 最长回文子串 https://leetcode.cn/problems/longest-palindromic-substring/

// 给你一个字符串 s，找到 s 中最长的回文子串。如果一个字符串从左向右写和从右向左写是一样的，这样的字符串就是回文串。

// 示例 1：
// 输入：s = "babad"
// 输出："bab" 或 "aba"

// 示例 2：
// 输入：s = "cbbd"
// 输出："bb"

package leetcode

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := expandAroundCenter(s, i, i)   // 对应奇数长度的回文串
		l2, r2 := expandAroundCenter(s, i, i+1) // 对应偶数长度的回文串
		// 比较两种情况下的回文串长度, 取最长的
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]
}

// 中心扩展法
func expandAroundCenter(s string, left, right int) (int, int) {
	// 从中心向两边扩展
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// 返回符合条件的子串的左右索引, 逆操作 left--, right++
	return left + 1, right - 1
}
