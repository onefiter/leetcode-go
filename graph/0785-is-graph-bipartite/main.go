package bipartite

import "leetcode-go/template"

// 785. 判断二分图
// https://leetcode.cn/problems/is-graph-bipartite/
// 给一个 n 个节点的无向图邻接表 graph（可能不连通），
// 判断能否把节点分成两个集合，使每条边的两端分属不同集合。
//
// 核心思路：二分图 ⇔ 可以二染色 ⇔ 不存在奇环。
// 对每个连通块 BFS 染色，相邻点颜色相同即失败；
// 并查集视角：u 的所有邻居必须同属一个集合，且不能与 u 同集合。

// isBipartite 解法一：BFS 二染色 ⭐ 最优方案
// 时间复杂度 O(n + E)，空间复杂度 O(n)。
func isBipartite(graph [][]int) bool {
	n := len(graph)
	color := make([]int, n) // 0 未染色，1 / -1 两种颜色
	for s := 0; s < n; s++ {
		if color[s] != 0 {
			continue
		}
		color[s] = 1
		q := []int{s}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v := range graph[u] {
				if color[v] == 0 {
					color[v] = -color[u]
					q = append(q, v)
				} else if color[v] == color[u] {
					return false
				}
			}
		}
	}
	return true
}

// isBipartiteUnionFind 解法二：并查集（对照用）
// 把 u 的所有邻居合并到同一集合（它们必须同色），若某邻居已与 u 连通则冲突。
// 时间复杂度 O((n + E) · α(n))，空间复杂度 O(n)。
func isBipartiteUnionFind(graph [][]int) bool {
	uf := template.NewUnionFind(len(graph))
	for u, nbs := range graph {
		for _, v := range nbs {
			if uf.Connected(u, v) {
				return false
			}
			uf.Union(nbs[0], v)
		}
	}
	return true
}
