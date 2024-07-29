// 905. 按奇偶排序数组 https://leetcode.cn/problems/sort-array-by-parity/

// 给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。
// 返回满足此条件的 任一数组 作为答案。

// 示例 1：
// 输入：nums = [3,1,2,4]
// 输出：[2,4,3,1]

// 示例 2：
// 输入：nums = [0]
// 输出：[0]

package leetcode

// 双指针
func sortArrayByParity(nums []int) []int {
	i, j := 0, len(nums)-1
	// 偶数在前，奇数在后
	for i < j {
		switch nums[i] & 1 {
		case 0: // 偶数
			i++
		case 1: // 奇数
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
	return nums
}
