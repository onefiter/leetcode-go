package trie

// 208. 实现 Trie (前缀树)
// https://leetcode.cn/problems/implement-trie-prefix-tree/
// 实现 Trie 类：Insert(word) 插入字符串，Search(word) 判断 word 是否被插入过，
// StartsWith(prefix) 判断是否有已插入的字符串以 prefix 为前缀。word 和 prefix 仅由小写字母组成。
//
// 核心思路：每个节点用长度 26 的子节点数组表示下一个字符，isEnd 标记「到这里是一个完整单词」。
// 插入时沿路径不存在就新建；查找时沿路径走，走不通返回 false。
// Search 要求终点 isEnd 为 true，StartsWith 只要路径走得通即可。

// Trie 解法一：26 叉数组前缀树 ⭐ 最优方案
// 时间复杂度：每个操作 O(L)，L 为字符串长度；空间复杂度 O(总字符数 * 26)。
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = &Trie{}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

// searchPrefix 沿 prefix 走到对应节点，走不通返回 nil。
func (this *Trie) searchPrefix(prefix string) *Trie {
	node := this
	for i := 0; i < len(prefix); i++ {
		node = node.children[prefix[i]-'a']
		if node == nil {
			return nil
		}
	}
	return node
}

// trieMapNode 用 map 存子节点的前缀树节点。
type trieMapNode struct {
	children map[byte]*trieMapNode
	isEnd    bool
}

// TrieMap 解法二：map 子节点前缀树（对照用）
// 不限于小写字母，稀疏时更省空间，但 map 的常数开销比数组大。
// 时间复杂度：每个操作 O(L)（均摊），空间复杂度 O(总字符数)。
type TrieMap struct {
	root *trieMapNode
}

func NewTrieMap() TrieMap {
	return TrieMap{root: &trieMapNode{children: map[byte]*trieMapNode{}}}
}

func (this *TrieMap) Insert(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		next, ok := node.children[word[i]]
		if !ok {
			next = &trieMapNode{children: map[byte]*trieMapNode{}}
			node.children[word[i]] = next
		}
		node = next
	}
	node.isEnd = true
}

func (this *TrieMap) Search(word string) bool {
	node := this.find(word)
	return node != nil && node.isEnd
}

func (this *TrieMap) StartsWith(prefix string) bool {
	return this.find(prefix) != nil
}

func (this *TrieMap) find(s string) *trieMapNode {
	node := this.root
	for i := 0; i < len(s); i++ {
		next, ok := node.children[s[i]]
		if !ok {
			return nil
		}
		node = next
	}
	return node
}
