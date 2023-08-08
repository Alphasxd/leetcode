// 41. 缺失的第一个正数 https://leetcode-cn.com/problems/first-missing-positive/

// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

// 示例 1：
// 输入：nums = [1,2,0]
// 输出：3

// 示例 2：
// 输入：nums = [3,4,-1,1]
// 输出：2

// 示例 3：
// 输入：nums = [7,8,9,11,12]
// 输出：1

package leetcode

func firstMissingPositive(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		if num > 0 {
			m[num] = true
		}
	}
	for i := 1; i <= len(nums); i++ {
		if !m[i] {
			return i
		}
	}
	// 如果 [1, len(nums)] 都出现了，那么答案就是 len(nums)+1
	return len(nums) + 1
}
