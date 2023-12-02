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

// 贪心算法
func maxProfit(prices []int) int {
	// 如果数组长度小于2，直接返回0，因为没有卖出的机会
	if len(prices) < 2 {
		return 0
	}

    maxPro, minPrice := 0, prices[0]

    // 遍历数组，记录股票价格最低点和卖出所得最大利润
	for _, price := range prices {
		// 如果当前值小于最小值，更新最小值
		minPrice = min(minPrice, price)
		// 如果当前值减去最小值大于最大利润，更新最大利润
		maxPro = max(maxPro, price-minPrice)
	}
	return maxPro
}