// 238. 除自身以外数组的乘积 https://leetcode.cn/problems/product-of-array-except-self/

// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
// 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

// 示例 1：
// 输入：nums = [1,2,3,4]
// 输出：[24,12,8,6]

// 示例 2：
// 输入：nums = [-1,1,0,-3,3]
// 输出：[0,0,9,0,0]

package leetcode

func productExceptSelf(nums []int) []int {
    length := len(nums)

    ans := make([]int, length)

    product := 1
    for i := range ans {
        ans[i] = product
        // 将 ans[i] 设为左侧所有元素的乘积
        product *= nums[i]
    }

    product = 1
    for i := length - 1; i >= 0; i-- {
        // 将上一步的结果乘以右侧所有元素的乘积即为答案
        ans[i] *= product
        // 将 ans[i] 设为 ans[i] * 右侧所有元素的乘积
        product *= nums[i]
    }

    return ans
}
