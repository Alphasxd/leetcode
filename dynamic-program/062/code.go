// 62. 不同路径 https://leetcode-cn.com/problems/unique-paths/

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
// 问总共有多少条不同的路径？

// 示例 1：
// 输入：m = 3, n = 7
// 输出：28

package leetcode

import "math/big"

// 动态规划
// 我们只需要一个长度为 n 的一维数组 dp，其中 dp[j] 存储的是到达当前行第 j 列的路径数量。
// 然后通过更新 dp[j] 的值来实现动态规划, 
// dp[j] 的新值等于 dp[j]（上边的路径数量, 同一列）和 dp[j-1]（左边列的路径数量）的和。
func uniquePaths(m, n int) int {
	dp := make([]int, n)
	// 到达第一行的任何位置只有一种走法, 所以初始化为 1
	for i := range dp {
		dp[i] = 1
	}
	// 从第二行开始，每个位置的走法等于左边(dp[j-1])和上边(dp[j])的走法之和
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	// 最后一个元素的值即为到达右下角的走法总数
	return dp[n-1]
}

// 实际上和杨辉三角有关，可以用排列组合的方法解决
func uniquePaths1(m, n int) int {
	// 从左上角到右下角，一共需要向右走 m-1 步，向下走 n-1 步
	// 一共需要走 m+n-2 步，其中 m-1 步向右，n-1 步向下
	// 从 m+n-2 中选择 m-1 步向右，共有 C(m+n-2, m-1) 种走法
	return int(new(big.Int).Binomial(int64(m+n-2), int64(m-1)).Int64())
}