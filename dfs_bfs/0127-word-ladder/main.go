package wordladder

// 127. 单词接龙
// https://leetcode.cn/problems/word-ladder/
// 给定 beginWord、endWord 和字典 wordList，每次只能改变一个字母且中间单词必须在字典中，
// 返回从 beginWord 到 endWord 的最短转换序列中的单词数目，不存在则返回 0。
//
// 核心思路：把每个单词看成图中的点，相差一个字母的单词之间连边，求最短路径 → BFS。
// 双向 BFS 从两端同时扩展，每次挑较小的一侧扩展，搜索空间从 b^d 降到 2*b^(d/2)。

// ladderLength 解法一：双向 BFS ⭐ 最优方案
// 时间复杂度 O(N*L*26)，空间复杂度 O(N*L)，N 为字典大小，L 为单词长度。
func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := map[string]bool{}
	for _, w := range wordList {
		dict[w] = true
	}
	if !dict[endWord] {
		return 0
	}
	front := map[string]bool{beginWord: true}
	back := map[string]bool{endWord: true}
	visited := map[string]bool{beginWord: true, endWord: true}
	for step := 1; len(front) > 0 && len(back) > 0; step++ {
		// 始终扩展较小的一侧
		if len(front) > len(back) {
			front, back = back, front
		}
		next := map[string]bool{}
		for w := range front {
			bs := []byte(w)
			for i := range bs {
				orig := bs[i]
				for ch := byte('a'); ch <= 'z'; ch++ {
					if ch == orig {
						continue
					}
					bs[i] = ch
					nw := string(bs)
					if back[nw] {
						// 两侧相遇
						return step + 1
					}
					if dict[nw] && !visited[nw] {
						visited[nw] = true
						next[nw] = true
					}
				}
				bs[i] = orig
			}
		}
		front = next
	}
	return 0
}

// ladderLengthBFS 解法二：单向 BFS（对照用）
// 从 beginWord 逐层扩展，直到遇到 endWord。
// 时间复杂度 O(N*L*26)，空间复杂度 O(N*L)，但实际扩展的节点数远多于双向 BFS。
func ladderLengthBFS(beginWord string, endWord string, wordList []string) int {
	dict := map[string]bool{}
	for _, w := range wordList {
		dict[w] = true
	}
	if !dict[endWord] {
		return 0
	}
	visited := map[string]bool{beginWord: true}
	queue := []string{beginWord}
	for step := 1; len(queue) > 0; step++ {
		var next []string
		for _, w := range queue {
			if w == endWord {
				return step
			}
			bs := []byte(w)
			for i := range bs {
				orig := bs[i]
				for ch := byte('a'); ch <= 'z'; ch++ {
					bs[i] = ch
					nw := string(bs)
					if dict[nw] && !visited[nw] {
						visited[nw] = true
						next = append(next, nw)
					}
				}
				bs[i] = orig
			}
		}
		queue = next
	}
	return 0
}
