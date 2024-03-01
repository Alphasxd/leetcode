package leetcode

func insertionSort(nums []int) {
	// 从第二个元素开始，将其插入到前面已经排好序的序列中
	for i := 1; i < len(nums); i++ {
		// 保存当前元素
		key := nums[i]
		// 从当前元素的前一个元素开始，依次向前比较，如果比当前元素大，则将其后移一位
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		// 此时 nums[j] <= key，将 key 插入到 j+1 的位置
		nums[j+1] = key
	}
}