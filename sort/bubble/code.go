package leetcode

// 冒泡排序
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		// 因为冒泡排序每次都会把最大的数放到最后，所以每次循环都可以减少一次循环
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
