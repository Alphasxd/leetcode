// 172. 阶乘后的零 https://leetcode.cn/problems/factorial-trailing-zeroes/

// 给定一个整数 n，返回 n! 结果尾数中零的数量。

package leetcode

// 阶乘后的零的数量取决于阶乘中存在多少个10，而10可分解为2和5，而2的数量又远大于5的数量，因此只需计算5的数量即可。
func trailingZeroes(n int) int {
	cnt := 0
	for n > 0 {
		cnt += n / 5
		n /= 5
	}
	return cnt
}
