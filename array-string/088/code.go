// 88. 合并两个有序数组 https://leetcode.cn/problems/merge-sorted-array/

// 给你两个非递减顺序排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n，分别表示 nums1 和 nums2 中的元素数目。
// 请你将 nums2 合并到 nums1 中，使合并后的数组同样按非递减顺序排列。
// nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

// 示例 1：
// 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// 输出：[1,2,2,3,5,6]

package leetcode

// 时间复杂度：O(m+n)，空间复杂度：O(1)
// 比较 nums1[i] 和 nums2[j]，将较大的元素放在 nums1[k]，然后移动指针。
// 重复此过程直到 nums2 的所有元素被合并。
func merge(nums1 []int, m int, nums2 []int, n int) {
    i, j, k := m-1, n-1, m+n-1
    for j >= 0 {
        if i >= 0 && nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
        k--
    }
}
