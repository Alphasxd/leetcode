// 1143. 最长公共子序列 https://leetcode.cn/problems/longest-common-subsequence/

// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。如果不存在公共子序列，返回 0 。
// 子序列: 是由原字符串删除一些(或不删除)字符而不改变剩余字符相对位置形成的新字符串

// 示例 1:
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3

// 示例 2:
// 输入：text1 = "abc", text2 = "def"
// 输出：0

package leetcode

func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	// dp[i][j] 表示 text1[0:i] 和 text2[0:j] 的最长公共子序列长度
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				// 如果 text1[i-1] == text2[j-1]，则 text1[i-1] 和 text2[j-1] 必然在最长公共子序列中
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 否则，text1[i-1] 和 text2[j-1] 至少有一个不在最长公共子序列中
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
