// 69. 加一 https://leetcode.cn/problems/plus-one/

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
// 你可以假设除了整数 0 之外，这个整数不会以零开头。

package leetcode

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			// 否则直接加一然后返回
			digits[i]++
			return digits
		}
	}
	// 如果循环结束还没有返回，说明数组中所有的元素都是9
	// 此时需要扩展数组长度，然后把首位置为1即可
	return append([]int{1}, digits...)
}
