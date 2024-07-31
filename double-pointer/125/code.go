// 125. 验证回文串 https://leetcode.cn/problems/valid-palindrome/

// 给你一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。
// 是返回 true ，否则返回 false 。

package leetcode

import (
    "strings"
    "unicode"
)

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    // 双指针
    i, j := 0, len(s)-1

    for i < j {
        // 非字母数字，跳过
        if !isAlphanumeric(rune(s[i])) {
            i++
            continue
        }
        if !isAlphanumeric(rune(s[j])) {
            j--
            continue
        }
        // 不相等，返回 false
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func isAlphanumeric(c rune) bool {
    return unicode.IsLetter(c) || unicode.IsDigit(c)
}
