// 287. 寻找重复数 https://leetcode.cn/problems/find-the-duplicate-number/

// 给定一个包含 n + 1 个整数的数组 nums，其数字都在 [1, n] 范围内，可知至少存在一个重复的整数。
// 假设 nums 只有一个重复的整数，找出这个重复的数。不得修改原数组（假设数组是只读的）。且只能使用额外的 O(1) 的空间。

// 示例1:
// 输入: [1,3,4,2,2]
// 输出: 2

// 示例2:
// 输入: [3,1,3,4,2]
// 输出: 3

package leetcode

// 因为 nums 中的数字范围在 [1, n] 之间，所以可以将 nums 中的每个数字看作一个指针，指向 nums 中的一个位置
func findDuplicate(nums []int) int {
	// slow, fast 分别指向 nums[0] 和 nums[nums[0]]
	slow := nums[nums[0]]
	fast := nums[nums[nums[0]]]
	// slow 一次走一步，fast 一次走两步，直到相遇
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// slow 从 nums[0] 开始，fast 从相遇点开始，再次相遇时即为重复元素
	duplicate := nums[0]
	for duplicate != slow {
		duplicate = nums[duplicate]
		slow = nums[slow]
	}
	return duplicate
}
