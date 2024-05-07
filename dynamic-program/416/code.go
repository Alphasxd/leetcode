// 416. 分割等和子集 https://leetcode-cn.com/problems/partition-equal-subset-sum/

// 给你一个只包含正整数的非空数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

// 示例 1：
// 输入：nums = [1, 5, 11, 5]
// 输出：true

// 示例 2：
// 输入：nums = [1, 2, 3, 5]
// 输出：false

package leetcode

// 如果数组的和是奇数，那么不可能分割成两个和相等的子集
// 如果数组的和是偶数，那么问题转化成背包问题，背包容量是数组和的一半
func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}

	target := sum >> 1
	// dp 表示是否可以从数组中选取若干个数，使得这些数的和等于 target
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[target]
}