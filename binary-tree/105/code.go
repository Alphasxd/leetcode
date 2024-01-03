// 从前序与中序遍历序列构造二叉树 https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

// 示例 1:
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
//              3
//             / \
//            9  20
//              /  \
//             15   7

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归, 时间复杂度O(n), 空间复杂度O(n)
func buildTree(preorder, inorder []int) *TreeNode {
	// 如果前序遍历为空, 则返回nil
	if len(preorder) == 0 {
		return nil
	}

	// 从中序遍历中找到根节点的索引，根节点即为前序遍历的第一个元素
	i := func(order []int, v int) int {
		var index int
		for order[index] != v {
			index++
		}
		return index
	}(inorder, preorder[0])

	// 递归构造左右子树
	// 根节点为 preorder[0], 中序遍历的根节点索引为 i
	// 左右子树的节点个数对应相等即可，因此可以得出 preorder在左右子树的划分
	// 前序遍历的左子树为 preorder[1:i+1], 中序遍历的左子树为 inorder[:i]
	// 前序遍历的右子树为 preorder[i+1:], 中序遍历的右子树为 inorder[i+1:]
	return &TreeNode{
		Val: preorder[0],
		Left: buildTree(preorder[1:i+1], inorder[:i]),
		Right: buildTree(preorder[i+1:], inorder[i+1:]),
	}
}
