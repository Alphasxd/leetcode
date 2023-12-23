// 290. 单词规律 https://leetcode-cn.com/problems/word-pattern/

// 给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。
// 这里的 遵循 指完全匹配，例如 pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。

// 示例1:
// 输入: pattern = "abba", s = "dog cat cat dog"
// 输出: true

// 示例 2:
// 输入:pattern = "abba", s = "dog cat cat fish"
// 输出: false

// 示例 3:
// 输入: pattern = "aaaa", s = "dog cat cat dog"
// 输出: false

package leetcode

import "strings"

// 与205题类似，只是这里的s是单词，而不是单个字符，只需要将 s 中的每个单词当作一个字符即可
func wordPattern(pattern string, s string) bool {
	// 首先分割字符串，取出每个单词
	words := strings.Fields(s)
	// 如果单词数量和pattern长度不一致，直接返回false
	if len(pattern) != len(words) {
		return false
	}
	// 两个map，分别记录pattern到单词的映射和单词到pattern的映射
	m1, m2 := make(map[byte]string), make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		// 如果两个map中都没有记录，则添加记录
		if m1[pattern[i]] == "" && m2[words[i]] == 0 {
			m1[pattern[i]] = words[i]
			m2[words[i]] = pattern[i]
		} else if m1[pattern[i]] != words[i] || m2[words[i]] != pattern[i] {
			// 如果两个map中有记录，但是不对应，则返回false
			return false
		}
	}
	return true
}