// 78. 子集 https://leetcode.cn/problems/subsets/

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// 示例 2：
// 输入：nums = [0]
// 输出：[[],[0]]

package leetcode

// 时间复杂度：O(n*2^n)
func subsets(nums []int) [][]int {
    // 创建一个 length 为 1，capicity 为 2^len(nums) 的二维数组，用于存储子集
    // length 为 1，正好将空集存储进去， 初始化后的 sets 为 [[]]
    sets := make([][]int, 1, 1<<uint(len(nums)))
    for _, num := range nums {
        for _, set := range sets {
            // 创建一个长度为 len(set)，capicity 为 len(set)+1 的切片
            s := make([]int, len(set), len(set)+1)
            // 将 set 中的元素复制到 s 中
            copy(s, set)
            // 首先将 num 添加到 s 中，然后将 s 添加到 sets 中
            sets = append(sets, append(s, num))
        }
    }
    return sets
}
