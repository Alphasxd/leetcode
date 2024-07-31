// 162. 寻找峰值 https://leetcode.cn/problems/find-peak-element/

// 峰值元素是指其值大于左右相邻值的元素。返回数组中的任何一个峰值即可

// 示例 1：
// 输入：nums = [1,2,3,1]
// 输出：2

// 示例 2：
// 输入：nums = [1,2,1,3,5,6,4]
// 输出：1 或 5

package leetcode

// 二分查找
func findPeakElement(nums []int) int {
    l, r := 0, len(nums)-1
    for l < r {
        h := int(uint(l+r) >> 1)
        if nums[h] > nums[h+1] {
            r = h
        } else {
            l = h + 1
        }
    }
    return l
}

// func findPeakElement(nums []int) int {
//     idx := 0
//     for i, v := range nums {
//         // 如果当前值比前一个值大，则继续向后搜索
//         if v > nums[idx] {
//             // 如果 v < nums[idx]，则不会更新 idx
//             idx = i
//         }
//     }
//     return idx
// }
