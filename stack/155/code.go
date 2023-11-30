// 155. 最小栈 https://leetcode-cn.com/problems/min-stack/

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

package leetcode

import "math"

type MinStack struct {
	stack []int
	minstack []int // 辅助存储最小元素的栈
}

// 初始化堆栈对象
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minstack: []int{math.MaxInt64}, // 初始化时将最大值赋值给辅助栈
	}
}

// 将元素 val 推入堆栈
// 当一个元素要入栈时，首先取当前辅助栈的栈顶存储的最小值，然后与当前元素比较出最小值
// 将得出的最小值推入辅助栈的栈顶，即辅助栈的栈顶存储当前栈所对应的最小值
func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack, val)
	top := this.minstack[len(this.minstack)-1]
	this.minstack = append(this.minstack, min(val, top))
}

// 删除堆栈顶部的元素，将辅助栈的栈顶元素一同弹出
func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.minstack = this.minstack[:len(this.minstack)-1]
}

// 获取堆栈顶部的元素
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

// 获取堆栈中的最小元素
func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack)-1]
}