// 121. 买卖股票的最佳时机 https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

// 给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。
// 设计一个算法来计算你所能获取的最大利润。返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0。

// 示例 1：
// 输入：[7,1,5,3,6,4]
// 输出：5

// 示例 2：
// 输入：prices = [7,6,4,3,1]
// 输出：0

package leetcode

func maxProfit(prices []int) int {
	// 如果数组长度小于2，直接返回0
	if len(prices) < 2 {
		return 0
	}

	maxPro := 0
	minPrice := prices[0]

	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxPro = max(maxPro, price-minPrice)
	}
	return maxPro
}