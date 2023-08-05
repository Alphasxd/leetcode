// 53. 最大子数组和 https://leetcode-cn.com/problems/maximum-subarray/

// 示例1：
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6

// 示例2：
// 输入：nums = [1]
// 输出：1

// 示例3：
// 输入：nums = [5,4,-1,7,8]
// 输出：23

package leetcode

// 动态规划
// func maxSubArray(nums []int) int {
// 	max := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i-1] + nums[i] > nums[i] {
// 			nums[i] += nums[i-1]
// 		}
// 		if nums[i] > max {
// 			max = nums[i]
// 		}
// 	}
// 	return max
// }

// 贪心算法
func maxSubArray(nums []int) int {
	var max, sum int
	for i, num := range nums {
		// sum 为当前遍历到的元素和
		sum += num
		// 如果 sum 大于 max 或者是第一次遍历时，将 sum 赋值给 max
		if sum > max || i == 0 {
			max = sum
		}
		// 如果 sum 小于 0，将 sum 置为 0，因为 0 > 负数
		if sum < 0 {
			sum = 0
		}
	}
	return max
}