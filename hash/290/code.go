// 290. 单词规律 https://leetcode.cn/problems/word-pattern/

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
func wordPattern(pattern, s string) bool {
    words := strings.Fields(s)
    if len(pattern) != len(words) {
        return false
    }
    m1, m2 := make(map[byte]string), make(map[string]byte)
    for i := range pattern {
        x, y := pattern[i], words[i]
        if m1[x] != "" && m1[x] != y || m2[y] > 0 && m2[y] != x {
            return false
        }
        m1[x] = y
        m2[y] = x
    }
    return true
}
