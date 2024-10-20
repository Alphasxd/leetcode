// 113. 路径总和 II https://leetcode-cn.com/problems/path-sum-ii

// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
// 叶子节点 是指没有子节点的节点。

package leetcode

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
    if root == nil {
        return nil
    }

	// 只存在根节点，判断根节点的值是否等于 targetSum
    if root.Left == nil && root.Right == nil {
        if root.Val != targetSum {
            return nil
        }
        return [][]int{{root.Val}}
    }

    var res [][]int
    targetSum -= root.Val

    // 辅助函数，用于处理子树
    process := func(subtree *TreeNode) {
        for _, path := range pathSum(subtree, targetSum) {
            // 将当前节点的值添加到路径的开头，这里切片的操作会消耗一定的时间
            path = append([]int{root.Val}, path...)
            res = append(res, path)
        }
    }

    // 处理左子树和右子树
    process(root.Left)
    process(root.Right)

    return res
}

func pathSum1(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var curPath []int
	var dfs func(node *TreeNode, curSum int)

	dfs = func(node *TreeNode, curSum int) {
		if node == nil {
			return
		}

		curPath = append(curPath, node.Val)
		curSum += node.Val

		if node.Left == nil && node.Right == nil && curSum == targetSum {
            // 这里需要拷贝一份，因为 curPath 是一个切片，后续会修改
			copyPath := make([]int, len(curPath))
			copy(copyPath, curPath)
			res = append(res, copyPath)
		}

        dfs(node.Left, curSum)
		dfs(node.Right, curSum)

		curPath = curPath[:len(curPath)-1]
	}

	dfs(root, 0)
	return res
}