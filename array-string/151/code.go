// 151. 反转字符串中单词 https://leetcode-cn.com/problems/reverse-words-in-a-string/

// 给定一个字符串s，逐个翻转字符串中的每个单词的顺序
// 单词是由非空字符组成的字符串. s中使用至少一个空格将字符串中的单词分隔开
// 返回颠倒 s 中单词顺序且用单个空格将其余单词分隔开的字符串

// 示例 1：
// 输入：s = "the sky is blue"
// 输出："blue is sky the"

// 示例 2：
// 输入：s = "  hello world  "
// 输出："world hello"

package leetcode

import "strings"

func reverseWords(s string) string {
	// 将 s 分割成字符串数组
	ans := strings.Fields(s)
	// 将字符串数组反转
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	// 将字符串数组用空格拼接成字符串
	return strings.Join(ans, " ")
}