package aliendict

import "strings"

// 269. 火星词典
// https://leetcode.cn/problems/alien-dictionary/
// 给定一个按火星语字典序排好序的单词列表 words，推导出火星语字母的顺序并返回（任意一种合法顺序）。
// 若顺序不合法（存在矛盾，或长单词排在它的前缀之前），返回 ""。
//
// 核心思路：相邻两个单词第一个不同的字符 a≠b 给出一条约束 a -> b；
// 所有约束构成有向图，对出现过的字母求拓扑序即可。有环则无解。
// 特殊情况：若 w1 比 w2 长且 w2 是 w1 的前缀（如 "abc" 在 "ab" 之前），直接无解。

// buildGraph 建图：返回邻接表（去重）、出现过的字母集合；不合法时 ok=false。
func buildGraph(words []string) (graph [26][]int, present [26]bool, ok bool) {
	var hasEdge [26][26]bool
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			present[w[i]-'a'] = true
		}
	}
	for i := 0; i+1 < len(words); i++ {
		w1, w2 := words[i], words[i+1]
		j := 0
		for j < len(w1) && j < len(w2) && w1[j] == w2[j] {
			j++
		}
		if j == len(w1) || j == len(w2) {
			// 一个是另一个的前缀：只有短的在前才合法
			if len(w1) > len(w2) {
				return graph, present, false
			}
			continue
		}
		a, b := int(w1[j]-'a'), int(w2[j]-'a')
		if !hasEdge[a][b] {
			hasEdge[a][b] = true
			graph[a] = append(graph[a], b)
		}
	}
	return graph, present, true
}

// alienOrder 解法一：建图 + Kahn 拓扑排序 ⭐ 最优方案
// 设 C 为所有单词总长度，时间复杂度 O(C)，空间复杂度 O(1)（字母表固定 26，边至多 26*26）。
func alienOrder(words []string) string {
	graph, present, ok := buildGraph(words)
	if !ok {
		return ""
	}
	var indeg [26]int
	total := 0
	for u := 0; u < 26; u++ {
		if present[u] {
			total++
		}
		for _, v := range graph[u] {
			indeg[v]++
		}
	}
	queue := []int{}
	for u := 0; u < 26; u++ {
		if present[u] && indeg[u] == 0 {
			queue = append(queue, u)
		}
	}
	var sb strings.Builder
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		sb.WriteByte(byte('a' + u))
		for _, v := range graph[u] {
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if sb.Len() != total {
		return "" // 有环
	}
	return sb.String()
}

// alienOrderDFS 解法二：建图 + DFS 三色标记 + 后序反转（对照用）
// 时间复杂度 O(C)，空间复杂度 O(1)。
func alienOrderDFS(words []string) string {
	graph, present, ok := buildGraph(words)
	if !ok {
		return ""
	}
	var color [26]int // 0=未访问，1=递归栈中，2=已完成
	post := make([]byte, 0, 26)
	var dfs func(u int) bool // 返回 false 表示发现环
	dfs = func(u int) bool {
		color[u] = 1
		for _, v := range graph[u] {
			if color[v] == 1 || (color[v] == 0 && !dfs(v)) {
				return false
			}
		}
		color[u] = 2
		post = append(post, byte('a'+u))
		return true
	}
	for u := 0; u < 26; u++ {
		if present[u] && color[u] == 0 && !dfs(u) {
			return ""
		}
	}
	// 后序反转即拓扑序
	for i, j := 0, len(post)-1; i < j; i, j = i+1, j-1 {
		post[i], post[j] = post[j], post[i]
	}
	return string(post)
}
