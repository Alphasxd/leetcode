// 169. 多数元素 https://leetcode-cn.com/problems/majority-element/

// 给定一个大小为 n 的数组 nums，返回其中的多数元素。多数元素是指在数组中出现次数大于 n/2 的元素。
// 可以假定数组是非空的，并且给定的数组总是存在多数元素。

// 示例 1:
// 输入: [3,2,3]
// 输出: 3

// 示例 2:
// 输入: [2,2,1,1,1,2,2]
// 输出: 2

package leetcode

// Boyer-Moore 投票算法
func majorityElement(nums []int) int {
	// 定义 众数 和 统计数
	var major, cnt int
	for _, num := range nums {
		// 如果计数清零则重新选定当前num为major
		if cnt == 0 {
			major = num
		}
		// 如果当前num等于major则计数加一，否则减一
		if num == major {
			cnt++
		} else {
			cnt--
		}
	}
	return major
}

// map 计数
func majorityElement1(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] > len(nums)/2 {
			return num
		}
	}
	return 0
}