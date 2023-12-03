// 274. H指数 https://leetcode-cn.com/problems/h-index/

// 给你一个整数数组 citations，其中 citations[i] 表示研究者的第 i 篇论文的引用次数，计算并返回该研究者的 h 指数。
// 一名科研人员的 h 指数 是指他（她）至少发表了 h 篇至少 被引用 h 次的论文。

// 示例1：
// 输入：citations = [3,0,6,1,5]
// 输出：3

// 示例2：
// 输入：citations = [1,3,1]
// 输出：1

package leetcode

// 计数排序
func hIndex(citations []int) int {
	// 定义一个变量n，表示数组的长度，papers[i]表示引用次数为i的论文数量
	n := len(citations)
	// 引用次数大于等于n的论文数量
	papers := make([]int, n+1)
	// 遍历数组，统计引用次数
	for _, c := range citations {
		// 取引用次数和n中较小的那个，因为引用次数大于n的论文数量对于h指数没有意义
		papers[min(c, n)]++
	}
	// 定义一个变量k，表示当前的h指数，引用次数大于等于k的论文数量至少为k篇，初始化为n
	k := n
	// 遍历引用次数，从大到小累加引用次数，直到引用次数大于等于k
	for s := papers[n]; k > s; s += papers[k] {
		k--
	}
	// 返回h指数
	return k
}