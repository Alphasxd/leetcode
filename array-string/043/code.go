// 43. 字符串相乘 https://leetcode.cn/problems/multiply-strings/

// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

// 注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。

// 示例 1:
// 输入: num1 = "2", num2 = "3"
// 输出: "6"

// 示例 2:
// 输入: num1 = "123", num2 = "456"
// 输出: "56088"

package leetcode

import "strings"

func multiply(num1, num2 string) string {
	ints := make([]int, len(num1)+len(num2))
	// 数组 ints 保存乘积的每一位
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			ints[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}
	// 进位
	for i := len(ints) - 1; i > 0; i-- {
		ints[i-1] += ints[i] / 10
		ints[i] %= 10
	}
	return intsToString(ints)
}

// 将数组转换为字符串
func intsToString(ints []int) string {
	var i int
	for i < len(ints)-1 && ints[i] == 0 {
		i++
	}
	ints = ints[i:]

	var b strings.Builder
	b.Grow(len(ints))
	for _, i := range ints {
		b.WriteByte('0' + byte(i))
	}
	return b.String()
}
