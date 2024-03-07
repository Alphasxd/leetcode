// 136. 只出现一次的数字 https://leetcode-cn.com/problems/single-number/

// 给你一个非空数组，这个数组中有一个数字只出现了一次，其他的数字都出现了两次。找出这个只出现一次的数字。

// 示例 1：
// 输入: [2,2,1]
// 输出: 1

package leetcode

func singleNumber(nums []int) int {
	var single int
	for _, num := range nums {
		// 异或运算，相同的数异或结果为0，0和任何数异或结果为任何数
		single ^= num
	}
	return single
}