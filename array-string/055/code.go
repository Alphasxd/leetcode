// 55. 跳跃游戏 https://leetcode.cn/problems/jump-game/

// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

// 示例 1：
// 输入：nums = [2,3,1,1,4]
// 输出：true

package leetcode

func canJump(nums []int) bool {
	// 定义一个变量cover，表示当前能够覆盖的最远位置，index表示数组的最后一个下标
	cover, index := 0, len(nums)-1
	for i := 0; i <= cover; i++ {
		cover = max(cover, i+nums[i])
		// 如果cover大于等于index，说明可以到达最后一个下标，返回true
		if cover >= index {
			return true
		}
	}
	return false
}
