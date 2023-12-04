// 380. O(1)时间插入、删除和获取随机元素

// 实现RandomizedSet 类：
// RandomizedSet() 初始化 RandomizedSet 对象
// bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
// bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
// int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。

package leetcode

import "math/rand"

type RandomizedSet struct {
	nums    []int // 存储元素
	indices map[int]int // 存储元素的值到下标的映射
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	// 如果 indices 中存在 val，返回 false
	if _, ok := rs.indices[val]; ok {
		return false
	}
	// 将 val 添加到 indices 中，其值为 nums 的长度，正好是 val 在 nums 中的下标
	rs.indices[val] = len(rs.nums)
	// 将 val 添加到 nums 中
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	id, ok := rs.indices[val]
	// 如果 indices 中不存在 val，返回 false
	if !ok {
		return false
	}
	// 最后一个元素的下标
	last := len(rs.nums) - 1
	// 将最后一个元素挪到 val 的位置
	rs.nums[id] = rs.nums[last] // 更新 nums
	rs.indices[rs.nums[id]] = id // 更新 indices
	// 然后删除最后一个元素
	rs.nums = rs.nums[:last] // 使用切片删除最后一个元素
	delete(rs.indices, val) // 删除 indices 中 key为 val 的元素
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}
