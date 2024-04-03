// 152.乘积最大子数组 https://leetcode.cn/problems/maximum-product-subarray/

// 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
// 测试用例的答案是一个 32-位 整数。

package leetcode

func maxProduct(nums []int) int {
	product, res := 1, nums[0]
	for i := range len(nums) {
		product *= nums[i]
		res = max(res, product)
		if nums[i] == 0 {
			// 重置 product，即当前子数组的乘积
			product = 1
		}
	}
	// 两次遍历，第一次从左到右，第二次从右到左
	// 主要是应对数组中有奇数个负数的情况
	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		product *= nums[i]
		res = max(res, product)
		if nums[i] == 0 {
			product = 1
		}
	}
	return res
}
