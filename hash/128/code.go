// 128. 最长连续序列 https://leetcode-cn.com/problems/longest-consecutive-sequence/

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

// 示例 1：
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

// 示例 2：
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9

package leetcode

func longestConsecutive(nums []int) int {

	// 创建一个 map，key 为 nums 中的元素
	m := make(map[int]bool)
	// 将 nums 中存在的元素全部存入 map，value 为 true
	for _, num := range nums {
		m[num] = true
	}

	longest := 0 // 最长连续序列的长度

	// 遍历 map
	for num := range m {
		// num不应该存在前驱数，如果存在则跳过本次循环
		// 确保 num 是连续序列的第一个数
		if m[num-1] {
			continue // 跳过本次循环
		}
		length := 1
		// 如果 num+1 存在，说明 num+1 已经被遍历过了，计算连续序列的长度
		for m[num+1] {
			num++
			length++
		}
		if length > longest {
			longest = length
		}
	}

	return longest
}
