// 219. 存在重复元素 II https://leetcode-cn.com/problems/contains-duplicate-ii/

// 判断数组中是否存在两个不同的索引 i 和 j，使得 nums[i] = nums[j]，并且 i 和 j 的差的绝对值最大为 k

// 示例 1:
// 输入: nums = [1,2,3,1], k = 3
// 输出: true

// 示例 2:
// 输入: nums = [1,0,1,1], k = 1
// 输出: true

// 示例 3:
// 输入: nums = [1,2,3,1,2,3], k = 2
// 输出: false

package leetcode

// 哈希表 + 滑动窗口 和 217 题类似, 只是在判断是否存在重复元素时，需要维护一个大小为 k 的滑动窗口
func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]struct{})
	for i, num := range nums {
		// 首先维护滑动窗口，保证哈希表中的元素个数始终小于等于 k 个
		if i > k {
			delete(dict, nums[i-k-1])
		}
		// 此时在大小为k的滑动窗口中，如果存在相同的数值，则直接返回true
		if _, ok := dict[num]; ok {
			return true
		}
		dict[num] = struct{}{}
	}
	return false
}