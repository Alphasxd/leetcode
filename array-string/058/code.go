// 58. 最后一个单词的长度 https://leetcode.cn/problems/length-of-last-word/

// 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
// 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。

package leetcode

func lengthOfLastWord(s string) int {
	cnt := 0
	length := len(s)
	// 从后往前遍历字符串，如果当前字符不是空格，则计数器加一，否则如果计数器不为零，则跳出循环
	for i := length - 1; i >= 0; i-- {
		if s[i] != ' ' {
			cnt++
		} else if cnt != 0 {
			break
		}
	}
	return cnt
}
