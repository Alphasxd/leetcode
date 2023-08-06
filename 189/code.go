// 189. 轮转数组 https://leetcode-cn.com/problems/rotate-array/

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

// 示例1：
// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]

// 示例2：
// 输入: nums = [-1,-100,3,99], k = 2
// 输出: [3,99,-1,-100]

package leetcode

func rotate(nums []int, k int) {
	k %= len(nums)
	// 先反转整个数组，再反转前 k 个，再反转后面的
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

// reverse 反转数组
func reverse(nums []int) {
	// 只需要反转一半的数组，j-i-1 是为了避免奇数数组时中间的数被反转两次
	for i, j := 0, len(nums); i < j/2; i++ {
		nums[i], nums[j-i-1] = nums[j-i-1], nums[i]
	}
}