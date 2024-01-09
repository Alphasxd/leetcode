// 69 x 的平方根 https://leetcode-cn.com/problems/sqrtx/

// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

package leetcode

func mySqrt(x int) int {
	lo, hi := 0, x
	for lo < hi {
		mid := hi - (hi-lo)/2
		temp := mid * mid
		switch {
		case temp > x:
			hi = mid - 1
		case temp < x:
			lo = mid
		default:
			return mid
		}
	}
	return lo
}
