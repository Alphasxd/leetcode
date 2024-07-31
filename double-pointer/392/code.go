// 392. 判断子序列 https://leetcode.cn/problems/is-subsequence/

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

package leetcode

func isSubsequence(s, t string) bool {
    m, n := len(s), len(t)
    // 双指针，i 指向 s，j 指向 t
    i, j := 0, 0
    for i < m && j < n {
        // 如果相等，i 向后移动一位
        if s[i] == t[j] {
            i++
        }
        // 不管是否相等，t 都要向后移动一位
        j++
    }
    // 如果 i 移动到了 s 的末尾，说明 s 是 t 的子序列
    return i == m
}
