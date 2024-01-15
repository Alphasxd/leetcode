// 208. 实现 Trie (前缀树) https://leetcode-cn.com/problems/implement-trie-prefix-tree/

// 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

// 请你实现 Trie 类：
// Trie() 初始化前缀树对象。
// void insert(String word) 向前缀树中插入字符串 word 。
// boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
// boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

package leetcode

type Trie struct {
	child [26]*Trie // 指向子节点的指针数组
	isEnd bool      // 是否为字符串的结尾
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.child[ch] == nil {
			node.child[ch] = &Trie{}
		}
		// 指向子节点，继续遍历
		node = node.child[ch]
	}
	// 遍历完，将当前节点标记为字符串的结尾
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.child[ch] == nil {
			return nil
		}
		node = node.child[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	// node != nil 说明 word 的所有前缀都存在, 且最后一个前缀的 isEnd 为 true
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}