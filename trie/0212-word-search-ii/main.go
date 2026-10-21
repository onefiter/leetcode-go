package wordsearch

// 212. 单词搜索 II
// https://leetcode.cn/problems/word-search-ii/
// 给定 m x n 字符网格 board 和单词列表 words，返回所有同时在网格和列表中出现的单词。
// 单词必须由相邻（上下左右）单元格的字母构成，同一单词内一个单元格只能用一次。
//
// 核心思路：把所有单词建成前缀树，再从每个格子出发 DFS 回溯，
// 一边走网格一边走前缀树：当前路径不是任何单词的前缀就立即剪枝。
// 走到单词结尾节点就记录答案并把 word 清空（去重）；
// 子树全部被找完后把该子节点从父节点上摘掉，后续搜索不再进入（剪枝）。

type trieNode struct {
	children [26]*trieNode
	word     string // 非空表示以此节点结尾的完整单词
}

// findWords 解法一：前缀树 + 回溯 + 剪枝 ⭐ 最优方案
// 时间复杂度 O(m * n * 3^(L-1))，L 为最长单词长度（首步 4 个方向，之后不回头最多 3 个）；
// 所有单词共享一次网格遍历。空间复杂度 O(总字符数 * 26)。
func findWords(board [][]byte, words []string) []string {
	root := &trieNode{}
	for _, w := range words {
		node := root
		for i := 0; i < len(w); i++ {
			c := w[i] - 'a'
			if node.children[c] == nil {
				node.children[c] = &trieNode{}
			}
			node = node.children[c]
		}
		node.word = w
	}

	m, n := len(board), len(board[0])
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var ans []string
	var dfs func(parent *trieNode, i, j int)
	dfs = func(parent *trieNode, i, j int) {
		ch := board[i][j]
		node := parent.children[ch-'a']
		if node == nil {
			return // 当前路径不是任何单词的前缀，剪枝
		}
		if node.word != "" {
			ans = append(ans, node.word)
			node.word = "" // 去重：同一个单词只收集一次
		}
		board[i][j] = '#' // 标记已访问
		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && board[x][y] != '#' {
				dfs(node, x, y)
			}
		}
		board[i][j] = ch // 回溯
		if isEmpty(node) {
			parent.children[ch-'a'] = nil // 子树已找完，摘掉节点剪枝
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(root, i, j)
		}
	}
	return ans
}

// isEmpty 节点既不是单词结尾也没有子节点。
func isEmpty(node *trieNode) bool {
	if node.word != "" {
		return false
	}
	for _, c := range node.children {
		if c != nil {
			return false
		}
	}
	return true
}

// findWordsEach 解法二：对每个单词单独跑一遍「单词搜索」DFS（对照用）
// 时间复杂度 O(k * m * n * 3^L)，k 为单词数，单词之间无法共享前缀，大数据会超时。
// 空间复杂度 O(L) 递归栈。
func findWordsEach(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])
	var exist func(w string, k, i, j int) bool
	exist = func(w string, k, i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != w[k] {
			return false
		}
		if k == len(w)-1 {
			return true
		}
		ch := board[i][j]
		board[i][j] = '#'
		found := exist(w, k+1, i-1, j) || exist(w, k+1, i+1, j) ||
			exist(w, k+1, i, j-1) || exist(w, k+1, i, j+1)
		board[i][j] = ch
		return found
	}
	seen := map[string]bool{}
	var ans []string
	for _, w := range words {
		if seen[w] {
			continue
		}
		seen[w] = true
	search:
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if exist(w, 0, i, j) {
					ans = append(ans, w)
					break search
				}
			}
		}
	}
	return ans
}
