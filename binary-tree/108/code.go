// 108. 将有序数组转换为二叉搜索树 https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/

// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
// 高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

// 示例 1:
// 输入：nums = [-10,-3,0,5,9]
// 输出：[0,-3,9,-10,null,5]
// 解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案

// 示例 2:
// 输入：nums = [1,3]
// 输出：[3,1]

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历，总是选择中间位置左边的数字作为根节点，时间复杂度O(n)，空间复杂度O(logn)
func sortedArrayToBTS(nums []int) *TreeNode {
	// 如果数组为空，返回nil
	if len(nums) == 0 {
		return nil
	}
	// mid为数组中间的元素
	mid := len(nums) / 2
	return &TreeNode{
		// 中间元素作为根节点
		Val:   nums[mid],
		// 递归调用，将数组左边的元素构造左子树，右边的元素构造右子树
		Left:  sortedArrayToBTS(nums[:mid]),
		Right: sortedArrayToBTS(nums[mid+1:]),
	}
}
