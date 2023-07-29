// 11. 盛最多水的容器 https://leetcode-cn.com/problems/container-with-most-water/

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

package leetcode

// S(i,j)=min(h[i],h[j])×(j−i)
// 容纳的水量取决于两个因素：两根指针指向的数字中较小值与指针之间的距离
func maxArea(height []int) int {
	ans := 0
	// 双指针，从两边向中间移动，每次移动较小的那个指针
	// 如果移动较大的那个指针，那么容器的容量一定会减小
	i, j := 0, len(height)-1
	for i < j {
		if height[i] < height[j] {
			ans = max(ans, height[i]*(j-i))
			i++
		} else {
			ans = max(ans, height[j]*(j-i))
			j--
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
