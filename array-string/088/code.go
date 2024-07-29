// 88. 合并两个有序数组 https://leetcode.cn/problems/merge-sorted-array/

// 给你两个非递减顺序排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n，分别表示 nums1 和 nums2 中的元素数目。
// 请你将 nums2 合并到 nums1 中，使合并后的数组同样按非递减顺序排列。
// nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

// 示例 1：
// 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// 输出：[1,2,2,3,5,6]

package leetcode

import "sort"

// 首先合并，然后排序
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 将 nums2 合并到 nums1 中的后 n 个位置
	// nums1[m:] 表示 nums1 中从 m 开始的所有元素，切片操作是左闭右开的
	copy(nums1[m:], nums2)
	// 调用 sort 包中的排序方法，底层是快速排序和插入排序的结合
	sort.Ints(nums1)
}
