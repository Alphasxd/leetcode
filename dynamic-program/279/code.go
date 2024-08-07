// 279. 完全平方数 https://leetcode.cn/problems/perfect-squares/

// 给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。

// 示例 1：
// 输入：n = 12
// 输出：3
// 解释：12 = 4 + 4 + 4

package leetcode

import "math"

func numSquares(n int) int {
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        cnt := math.MaxInt32
        for j := 1; j*j <= i; j++ {
            cnt = min(cnt, dp[i-j*j])
        }
        dp[i] = cnt + 1
    }
    return dp[n]
}