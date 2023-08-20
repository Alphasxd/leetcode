// 283. 移动零 https://leetcode-cn.com/problems/move-zeroes/

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
	i, j := 0, 0
	for ; i < len(nums); i++ {
		// 如果当前元素不为0，就把当前元素放到j位置，然后j++
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	// 将 nums[j:] 全部置为0
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}
