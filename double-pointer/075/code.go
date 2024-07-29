// 75. 颜色分类 https://leetcode.cn/problems/sort-colors

// 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 用 0, 1, 2 分别表示红色、白色和蓝色。不得使用代码库中的排序函数来解决这道题。

// 示例 1：
// 输入：nums = [2,0,2,1,1,0]
// 输出：[0,0,1,1,2,2]

package leetcode

// 双指针，zero 指向 0 的最右边界，two 指向 2 的最左边界
func sortColors(nums []int) {
	zero, two := 0, len(nums)-1
	// one 从左往右遍历，遇到 0 则与 zero 交换，遇到 2 则与 two 交换
	for one := 0; one <= two; {
		switch nums[one] {
		case 0:
			nums[zero], nums[one] = nums[one], nums[zero]
			// zero 和 one 同步右移
			zero++
			one++
		case 1:
			one++
		case 2:
			nums[one], nums[two] = nums[two], nums[one]
			// two 左移，one 不动
			two--
		}
	}
}
