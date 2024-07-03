// 217. 存在重复元素 https://leetcode.cn/problems/contains-duplicate/description/

// 给你一个整数数组 nums
// 如果任一值在数组中出现至少两次，返回 true；如果数组中每个元素都不同，返回 false

// 示例 1：
// 输入：nums = [1,2,3,1]
// 输出：true

// 示例 2：
// 输入：nums = [1,2,3,4]
// 输出：false

package leetcode

func containsDuplicate(nums []int) bool {
	dict := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := dict[num]; ok {
			return true
		}
		dict[num] = struct{}{}
	}
	return false
}