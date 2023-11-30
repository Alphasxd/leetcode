// 80. 删除有序数组中的重复项 II https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 最多出现两次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

package leetcode

func removeDuplicates(nums []int) int {
	n := len(nums)
	// 当数组中的元素数量≤2时，直接返回len即可
	if n <= 2 {
		return n
	}
	// 定义双指针，初始都指向第三个元素，即下标2
	slow := 2
	for fast := 2; fast < n; fast++ {
		if nums[slow-2] != nums[fast]{
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}