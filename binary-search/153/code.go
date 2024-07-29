// 153. 寻找旋转排序数组中的最小值 https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/

// 给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

// 示例 1：
// 输入：nums = [3,4,5,1,2]
// 输出：1

package leetcode

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1

	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		// 如果中间值大于最右边的值，说明最小值在右边
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			// 包含 nums[mid] == nums[hi] 的情况，所以不能用 hi = mid - 1
			hi = mid
		}
	}
	return nums[lo]
}
