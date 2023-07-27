// 1. 两数之和 https://leetcode-cn.com/problems/two-sum/

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案

// 示例 1：
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]

// 示例 2：
// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]

package leetcode

func twoSum(nums []int, target int) []int {
	// 创建一个 map，key 为数组的值，value 为数组的索引
	m := make(map[int]int)
	// 遍历数组
	for i, num := range nums {
		// 判断 map 中是否存在 target - num 的值
		if j, ok := m[target-num]; ok {
			// 如果存在，返回 map 中的索引 j 和当前的索引 i
			return []int{j, i}
		}
		// 如果不存在，将当前的 num 存入 map 中
		m[num] = i
	}
	return nil
}