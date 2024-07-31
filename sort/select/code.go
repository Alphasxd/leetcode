package leetcode

// selectSort 对输入的整数切片 nums 进行选择排序
func selectSort(nums []int) {
    // 外层循环，遍历每一个元素，除了最后一个
    for i := 0; i < len(nums)-1; i++ {
        // 假设当前元素是最小的
        min := i
        // 内层循环，从 i+1 开始遍历剩余的元素
        for j := i + 1; j < len(nums); j++ {
            // 如果找到一个比当前最小元素还小的元素，更新最小元素的索引
            if nums[j] < nums[min] {
                min = j
            }
        }
        // 如果最小元素的索引不是 i，说明找到了一个比 nums[i] 更小的元素，交换它们
        if min != i {
            nums[i], nums[min] = nums[min], nums[i]
        }
    }
}
