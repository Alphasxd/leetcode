// 31. 下一个排列 https://leetcode.cn/problems/next-permutation

// 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。原地修改，只使用额外常数空间。

// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[1,3,2]

// 示例 2：
// 输入：nums = [3,2,1]
// 输出：[1,2,3]

package leetcode

// 可以直接将字典排序理解为一个整数比较大小的问题，找到下一个比当前大的最小的整数即可
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1

	// 找到 i < j
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		// 找到 i < k
		for nums[i] >= nums[k] {
			k--
		}
		// 交换 i 和 k
		nums[i], nums[k] = nums[k], nums[i]
	}
	// reverse [j:end]
	reverse(nums, j, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
