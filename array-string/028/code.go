// 28. 找出字符串中第一个匹配项的下标 https://leetcode.cn/problems/implement-strstr/

// 给你两个字符串 haystack 和 needle ， 请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（ 下标从 0 开始）。
// 如果 needle 不是 haystack 的一部分， 则返回 - 1 。

package leetcode

import "strings"

func strStr(haystack string, needle string) int {
    return strings.Index(haystack, needle)
}
