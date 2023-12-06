// 167. 两数之和 II - 输入有序数组 https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/

// 给你一个下标从 1 开始的整数数组 numbers，该数组已按非递减顺序排列。请你找出两个数满足相加之和等于目标数 target 。
// 以长度为 2 的整数数组返回你满足条件的下标，其中 index1 必须小于 index2 。

package leetcode

// 和 hash/001/code.go 中的 twoSum 函数相比，这里的数组是有序的，所以可以使用双指针法
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		// switch 更优雅
		switch {
		case sum < target:
			l++
		case sum > target:
			r--
		default:
			return []int{l + 1, r + 1}
		}
	}
	return nil
}
