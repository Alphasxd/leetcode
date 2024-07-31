// 72. 编辑距离 https://leetcode.cn/problems/edit-distance/

// 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
// 你可以对一个单词进行如下三种操作：
// 1. 插入一个字符 2. 删除一个字符 3. 替换一个字符

// 示例 1：
// 输入：word1 = "horse", word2 = "ros"
// 输出：3

// 示例 2：
// 输入：word1 = "intention", word2 = "execution"
// 输出：5

package leetcode

func minDistance(word1, word2 string) int {
    m, n := len(word1), len(word2)
    // dp[i][j] 表示 word1[:i] 转换成 word2[:j] 所使用的最少操作数
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    // 初始化
    for i := 0; i <= m; i++ {
        dp[i][0] = i
    }
    for j := 0; j <= n; j++ {
        dp[0][j] = j
    }

    // 动态规划
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            // 如果 word1[i-1] == word2[j-1]，则不需要操作
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                // 否则，取插入、删除、替换操作的最小值
                dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
            }
        }
    }
    return dp[m][n]
}
