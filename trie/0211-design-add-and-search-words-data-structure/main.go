package worddict

// 211. 添加与搜索单词 - 数据结构设计
// https://leetcode.cn/problems/design-add-and-search-words-data-structure/
// 设计 WordDictionary：AddWord(word) 添加单词，Search(word) 判断是否存在匹配的单词，
// word 中可以包含 '.'，'.' 可以匹配任意一个字母。单词仅由小写字母组成。
//
// 核心思路：前缀树 + DFS。AddWord 就是普通的 Trie 插入；
// Search 逐字符向下走，普通字母只走对应子节点，遇到 '.' 就枚举全部非空子节点递归，
// 任意一条分支匹配到末尾且 isEnd 为 true 即成功。

type trieNode struct {
	children [26]*trieNode
	isEnd    bool
}

// WordDictionary 解法一：前缀树 + DFS ⭐ 最优方案
// AddWord 时间复杂度 O(L)；Search 无 '.' 时 O(L)，最坏（全是 '.'）O(26^L)，但受树中实际节点数限制。
// 空间复杂度 O(总字符数 * 26)。
type WordDictionary struct {
	root *trieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &trieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = &trieNode{}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(node *trieNode, i int) bool
	dfs = func(node *trieNode, i int) bool {
		if i == len(word) {
			return node.isEnd
		}
		if word[i] == '.' {
			for _, child := range node.children {
				if child != nil && dfs(child, i+1) {
					return true
				}
			}
			return false
		}
		child := node.children[word[i]-'a']
		return child != nil && dfs(child, i+1)
	}
	return dfs(this.root, 0)
}

// WordDictionaryBrute 解法二：按长度分组 + 逐个比对（对照用）
// 把单词按长度存到 map[int][]string，Search 时只和同长度的单词逐字符比对。
// AddWord 时间复杂度 O(1)（均摊）；Search 时间复杂度 O(n * L)，n 为同长度单词数。
// 空间复杂度 O(总字符数)。
type WordDictionaryBrute struct {
	byLen map[int][]string
}

func NewWordDictionaryBrute() WordDictionaryBrute {
	return WordDictionaryBrute{byLen: map[int][]string{}}
}

func (this *WordDictionaryBrute) AddWord(word string) {
	this.byLen[len(word)] = append(this.byLen[len(word)], word)
}

func (this *WordDictionaryBrute) Search(word string) bool {
	for _, w := range this.byLen[len(word)] {
		if match(w, word) {
			return true
		}
	}
	return false
}

// match 判断等长的 w 是否匹配模式 pattern（'.' 匹配任意字母）。
func match(w, pattern string) bool {
	for i := 0; i < len(w); i++ {
		if pattern[i] != '.' && pattern[i] != w[i] {
			return false
		}
	}
	return true
}
