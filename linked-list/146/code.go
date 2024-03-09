// 146. LRU Cache https://leetcode-cn.com/problems/lru-cache/

// 请你设计并实现一个满足 LRU 缓存约束的数据结构
// 实现 LRUCache 类：
// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，
// 则插入该组「关键字-值」。当缓存容量达到上限时，则应该逐出最久未使用的关键字。
// 函数 get(key) 和 put(key, value) 的时间复杂度均为 O(1)

package leetcode

import "container/list"

// 键值对，实现 Value 接口，方便链表操作
type pair struct {
	key, value int
}

// LRU 缓存，使用双向链表和哈希表实现
type LRUCache struct {
	capacity int                   // 容量
	list     *list.List            // 双向链表
	cache    map[int]*list.Element // 哈希表
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity,
		list.New(),
		make(map[int]*list.Element),
	}
}

func (c *LRUCache) Get(key int) int {
	// 如果 key 存在，将节点移到链表头部，并返回节点的值
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		// 需要断言，因为 elem.Value 是 interface{} 类型
		return elem.Value.(pair).value
	}
	// key 不存在，返回 -1
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	// 如果 key 存在，更新节点的值，并将节点移到链表头部
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		// 更新对应 key 的 value
		elem.Value = pair{key, value}
		return
	}

	// 添加新节点之前，需要判断缓存是否已满
	// 如果 cache 已满，移除链表尾部节点，并删除哈希表中对应的项以及双向链表中的节点
	// 删除链表尾部节点就符合了 LRU 的要求，因为尾部节点是最久未使用的，每次插入新节点或者更新节点都是在链表头部
	if c.list.Len() == c.capacity {
		last := c.list.Back()
		delete(c.cache, last.Value.(pair).key)
		c.list.Remove(last)
	}
	// 在链表头部插入新节点，并在哈希表中添加 key 和节点的映射
	// PushFront 返回的是 *list.Element，所以可以直接赋值给 map，同时完成了两个操作
	c.cache[key] = c.list.PushFront(pair{key, value})
}
