// 70. 爬楼梯 https://leetcode.cn/problems/climbing-stairs/

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

// 示例 1：
// 输入： n = 2
// 输出： 2

// 示例 2：
// 输入： n = 3
// 输出： 3

package leetcode

// 思路：动态规划，dp[i] = dp[i-1] + dp[i-2]，dp[i] 表示爬到第 i 阶的方法数
// 爬到第i阶的方法数 = 爬到第i-1阶的方法数 + 爬到第i-2阶的方法数，因为每次可以爬1或2个台阶
func climbStairs(n int) int {
	// p:dp[i-2]，q:dp[i-1]，r:dp[i]
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
