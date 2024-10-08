// 41. 缺失的第一个正数 https://leetcode.cn/problems/first-missing-positive/

// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

// 示例 1：
// 输入：nums = [1,2,0]
// 输出：3

// 示例 2：
// 输入：nums = [3,4,-1,1]
// 输出：2

// 示例 3：
// 输入：nums = [7,8,9,11,12]
// 输出：1

package leetcode

// 集合，时间复杂度：O(n)，空间复杂度：O(n)
func firstMissingPositive1(nums []int) int {
    set := make(map[int]struct{}, len(nums))
    for _, v := range nums {
        if v > 0 {
            set[v] = struct{}{}
        }
    }
    for i := 1; i <= len(nums); i++ {
        if _, ok := set[i]; !ok {
            return i
        }
    }
    return len(nums) + 1
}

// 原地哈希，时间复杂度：O(n)，空间复杂度：O(1)
// 思路是自定义哈希函数，将值为i的数放到下标为i-1的位置上
// 让每一个 v 回到自己的位置 v-1 上，那些不符合的位置就是本该存在的数，所以最后返回不符合位置的下标+1
func firstMissingPositive(nums []int) int {
    // 自定义哈希，将值为 v 的数放到下标为 v-1 的位置上
    for _, v := range nums {
        for v > 0 && v <= len(nums) && nums[v-1] != v {
            nums[v-1], v = v, nums[v-1]
        }
    }

    for i := 1; i <= len(nums); i++ {
        if nums[i-1] != i {
            return i
        }
    }

    return len(nums) + 1
}
