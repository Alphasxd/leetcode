// 230. 二叉搜索树中第K小的元素 https://leetcode.cn/problems/kth-smallest-element-in-a-bst/

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第k个最小元素（从 1 开始计数）。

// 示例 1：
// 输入：root = [3,1,4,null,2], k = 1
//     3
//    / \
//   1   4
//    \
//     2
// 输出：1

// 示例 2：
// 输入：root = [5,3,6,2,4,null,null,1], k = 3
//            5
//           / \
//          3   6
//         / \
//        2   4
//       /
//      1
// 输出：3

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历，时间复杂度O(n)，空间复杂度O(n)
func kthSmallest(root *TreeNode, k int) int {
	// 存储中序遍历结果
	res := []int{}
	var inorder func(*TreeNode)
	inorder = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		inorder(tn.Left)
		res = append(res, tn.Val)
		inorder(tn.Right)
	}
	inorder(root)
	// 返回第k个最小元素
	return res[k-1]
}
