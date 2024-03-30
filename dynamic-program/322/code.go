// 零钱兑换 https://leetcode-cn.com/problems/coin-change/

// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
// 你可以认为每种硬币的数量是无限的。

// 示例 1：
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3

// 示例 2：
// 输入：coins = [2], amount = 3
// 输出：-1

package leetcode

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	// 初始化为 amount+1，不可能取到的值
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				// dp[i-coin]+1：使用一枚 coin 面值的硬币，然后计算剩余金额 i-coin 的最少硬币数
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}