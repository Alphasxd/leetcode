// 283. 移动零 https://leetcode.cn/problems/move-zeroes/

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

// 示例 1:
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]

// 示例 2:
// 输入: nums = [0]
// 输出: [0]

package leetcode

func moveZeroes(nums []int) {
	// 因为要保持非零元素的相对顺序，所以i和j都从0开始
	for i, j := 0, 0; j < len(nums); j++ {
		// j指向非零元素时，交换i和j的值，保证i之前的元素都是非零的
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
