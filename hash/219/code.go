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

// 哈希表 + 滑动窗口
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	// key: 数值 value: 下标索引 类似两数之和中的 map
// 	m := make(map[int]int)
// 	for i, num := range nums {
// 		// 如果存在相同的数值，判断下标索引是否满足条件
// 		// 此时的下标索引 i 是新索引，j 是旧索引，这样近似实现了滑动窗口
// 		if j, ok := m[num]; ok && i-j <= k {
// 			return true
// 		}
// 		m[num] = i
// 		// 如果哈希表中的元素个数大于 k 个，删除最旧的元素
// 		// 保证哈希表中的元素个数始终小于等于 k 个，节省空间开销
// 		if len(m) > k {
// 			delete(m, nums[i-k])
// 		}
// 	}
// 	return false
// }

// 进一步优化内存空间，map的定义改成 map[int]struct{}，map中的value不再存储任何信息，空结构体不占用内存空间
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]struct{})
	for i, num := range nums {
		// 首先维护滑动窗口，保证哈希表中的元素个数始终小于等于 k 个
		if i > k {
			delete(m, nums[i-k-1])
		}
		// 此时在大小为k的滑动窗口中，如果存在相同的数值，则直接返回true
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}