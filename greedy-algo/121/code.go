// 121. 买卖股票的最佳时机 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

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
    var minIndex, bonus int
    for i, p := range prices {
        // 首先列出 profit 的计算公式，当前价格减去最小价格
        profit := p - prices[minIndex]
        if profit > bonus {
            bonus = profit
        } else if profit < 0 { // 当前 profit < 0，也就是当前价格小于最小价格
            // 更新最小价格的索引
            minIndex = i
        }
    }
    return bonus
}
