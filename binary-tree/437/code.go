// 437. 路径总和 III https://leetcode.cn/problems/path-sum-iii/

// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的路径的数目。
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

// 示例 1：
// 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// 输出：3
// 解释：和等于 8 的路径有 3 条，如图所示。
//        10
//       /  \
//      5   -3
//     / \    \
//    3   2   11
//   / \   \
//  3  -2   1

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	// 使用字面量方式初始化map，key为前缀和，value为前缀和出现的次数
	preSumMap := map[int]int{0: 1}
	// 声明一个用于递归的函数变量
	var f func(*TreeNode, int) int
	f = func(root *TreeNode, curSum int) int {
		var ans int
		if root == nil {
			return 0
		}
		// 当前路径上的和
		curSum += root.Val
		// 如果前缀和中存在curSum-targetSum，说明存在一条路径和为targetSum
		// 这条路径即为curSum-(curSum-targetSum)=targetSum
		if cnt, ok := preSumMap[curSum-targetSum]; ok {
			ans += cnt
		}
		// 更新前缀和curSum的次数
		preSumMap[curSum]++

		// 递归左右子树
		ans += f(root.Left, curSum)
		ans += f(root.Right, curSum)

		// 回溯，恢复状态
		preSumMap[curSum]--
		return ans
	}
	return f(root, 0)
}
