// 448. 找到所有数组中消失的数字 https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/

// 给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。
// 请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

// 示例 1：
// 输入：nums = [4,3,2,7,8,2,3,1]
// 输出：[5,6]

package leetcode

// leetcode 41. 缺失的第一个正数，类似题目
// 字典，时间复杂度：O(n)，空间复杂度：O(n)
func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)
	dict := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		dict[v] = struct{}{}
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := dict[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}

// 原地哈希，时间复杂度：O(n)，空间复杂度：O(1)
func findDisappearedNumbers1(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		for v > 0 && v <= len(nums) && nums[v-1] != v {
			nums[v-1], v = v, nums[v-1]
		}
	}
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] != i {
			res = append(res, i)
		}
	}
	return res
}