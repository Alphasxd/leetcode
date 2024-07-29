// 198. 打家劫舍 https://leetcode.cn/problems/house-robber/

// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，
// 一夜之内能够偷窃到的最高金额。

// 示例 1：
// 输入：[1,2,3,1]
// 输出：4

// 示例 2：
// 输入：[2,7,9,3,1]
// 输出：12

package leetcode

func rob(nums []int) int {
	// 边界条件
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// 当房间数大于等于2
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		// first表示前i-1个房间的最大值，second表示前i个房间的最大值
		first, second = second, max(first+nums[i], second)
	}
	return second
}
