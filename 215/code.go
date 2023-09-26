// 215. 数组中的第K个最大元素 https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。请设计时间复杂度为 O(n) 的算法解决此问题。

// 示例 1:
// 输入: [3,2,1,5,6,4] 和 k = 2
// 输出: 5

// 示例 2:
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
// 输出: 4

package leetcode

// 快速选择思想，找到下标为 k-1 的元素即可，时间复杂度 O(n)
func findKthLargest(nums []int, k int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := partition(nums, lo, hi)
		switch {
		case k-1 < mid:
			hi = mid - 1
		case k-1 > mid:
			lo = mid + 1
		default:
			return nums[mid]
		}
	}
	// lo == hi
	return nums[lo]
}

// 一次快排，返回枢轴元素的下标
func partition(nums []int, lo, hi int) int {
	pivot := nums[lo]
	i, j := lo, hi
	for i < j {
		for i < j && nums[j] <= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] >= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	return i
}
