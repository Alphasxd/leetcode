// 228. 汇总区间 https://leetcode.cn/problems/summary-ranges/

// 给定一个 无重复元素 的有序整数数组 nums 。
// 返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。

package leetcode

import "strconv"

func summaryRanges(nums []int) []string {
	var res []string
	for i, n := 0, len(nums); i < n; {
		left := i
		// 从 i 开始，找到第一个不连续的位置，此时 i-1 为连续区间的最后一个位置
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		// 如果 left < i-1，格式化为 "left->right"，否则只有一个数
		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		res = append(res, s)
	}
	return res
}
