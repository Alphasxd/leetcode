// 45. 跳跃游戏Ⅱ https://leetcode.cn/problems/jump-game-ii/

// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
// 每个元素 nums[i] 表示从索引 i 能够跳到的最大长度。换句话说，如果你在 nums[i] 你可以跳 j ∈ [1, nums[i]] 步到达的位置。
// 请你返回到达数组最后一个位置的最少跳跃次数。

// 示例 1：
// 输入：nums = [2,3,1,1,4]
// 输出：2

// 示例 2：
// 输入：nums = [2,3,0,1,4]
// 输出：2

package leetcode

func jump(nums []int) int {
	// 定义一个变量steps，表示跳跃的次数，position表示当前能够到达的最远位置
	steps, position := 0, len(nums)-1
	for position > 0 {
		for i := 0; i < position; i++ {
			// 如果当前位置能够到达的最远位置大于等于position，说明从当前位置可以一步跳到最后一个位置
			if i+nums[i] >= position {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}
