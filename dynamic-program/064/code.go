// 64. 最小路径和 https://leetcode-cn.com/problems/minimum-path-sum/

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

// 示例 1：
// 输入：grid = [
// [1,3,1],
// [1,5,1],
// [4,2,1]
// ]
// 输出：7

package leetcode

// 和第 62 题类似，不过是格子加了权重，求最小路径和
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	// 初始化第一行
	dp[0] = grid[0][0]
	// 第一行中除了第一个元素外，每个元素的路径和等于左边的路径和加上当前格子的权重
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	// 从第二行开始，每个元素的路径和等于上边元素的路径和(dp[j])和左边元素的路径和(dp[j-1])的较小值加上当前格子的权重
	for i := 1; i < m; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
		}
	}
	return dp[n-1]
}