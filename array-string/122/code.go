// 122. 买卖股票的最佳时机 II https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

// 给你一个整数数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。
// 在每一天，你可以决定是否购买和售出股票。你在任何时候最多只能持有一股股票。
// 你也可以先购买，然后再同一天出售股票。返回你能获得的最大利润。

package leetcode

func maxProfit(prices []int) int {
	var ans int
	length := len(prices)
	for i := 1; i < length; i++ {
		profit := prices[i] - prices[i-1]
		// 只要今天的价格比昨天高，就卖出
		if profit > 0 {
			ans += profit
		}
	}
	return ans
}