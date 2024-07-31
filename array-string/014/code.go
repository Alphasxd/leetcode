// 14. 最长公共前缀 https://leetcode.cn/problems/longest-common-prefix/

// 编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回空字符串 ""。

package leetcode

import "strings"

// 使用 strings.HasPrefix() 函数判断字符串是否以指定前缀开头
func longestCommonPrefix(strs []string) string {
    // 如果字符串数组为空，则不存在公共前缀，直接返回空字符串
    if len(strs) == 0 {
        return ""
    }
    // 初始化公共前缀为第一个字符串
    prefix := strs[0]
    // 遍历字符串数组
    for _, str := range strs[1:] {
        // 如果当前字符串不是以公共前缀开头，则不断缩短公共前缀
        for !strings.HasPrefix(str, prefix) {
            // 如果公共前缀为空，则不存在公共前缀，直接返回空字符串
            if prefix == "" {
                return ""
            }
            // 缩短公共前缀
            prefix = prefix[:len(prefix)-1]
        }
    }
    return prefix
}
