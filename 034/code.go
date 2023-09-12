// 34. 在排序数组中查找元素的第一个和最后一个位置 https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]。
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

// 示例 1：
// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]

// 示例 2：
// 输入：nums = [5,7,7,8,8,10], target = 6
// 输出：[-1,-1]

// 示例 3：
// 输入：nums = [], target = 0
// 输出：[-1,-1]

package leetcode

import "sort"

// 二分查找，时间复杂度 O(logn)
func searchRange(nums []int, target int) []int {
	// 使用 SearchInts 在 nums 中搜索 target，如果找到则返回其索引，否则返回将会插入的索引
	left := sort.SearchInts(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	// 找到 target+1 的索引，再往前移动一个位置, 即为 target 的最右索引
	right := sort.SearchInts(nums, target+1) - 1
	return []int{left, right}
}