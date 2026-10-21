package redundant

import "leetcode-go/template"

// 684. 冗余连接
// https://leetcode.cn/problems/redundant-connection/
// 一棵 n 个节点（编号 1..n）的树添加了一条额外的边，得到图 edges。
// 找出一条可以删去的边，使剩余部分仍是树；若有多个答案，返回 edges 中最后出现的那条。
//
// 核心思路：按顺序加边，第一条「两端已经连通」的边一定在环上，
// 且由于环上其余边都更早出现，它就是环上最后出现的边，即答案。

// findRedundantConnection 解法一：并查集 ⭐ 最优方案
// 时间复杂度 O(n · α(n))，空间复杂度 O(n)。
func findRedundantConnection(edges [][]int) []int {
	uf := template.NewUnionFind(len(edges) + 1)
	for _, e := range edges {
		if !uf.Union(e[0], e[1]) {
			return e
		}
	}
	return nil
}

// findRedundantConnectionDFS 解法二：逐边 DFS 判连通（对照用）
// 加入每条边之前，先 DFS 检查 u 是否已能走到 v，能走到说明这条边成环。
// 时间复杂度 O(n²)，空间复杂度 O(n)。
func findRedundantConnectionDFS(edges [][]int) []int {
	n := len(edges)
	graph := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		visited := make([]bool, n+1)
		var dfs func(x int) bool
		dfs = func(x int) bool {
			if x == v {
				return true
			}
			visited[x] = true
			for _, y := range graph[x] {
				if !visited[y] && dfs(y) {
					return true
				}
			}
			return false
		}
		// 两端都已在图中才可能已连通
		if len(graph[u]) > 0 && len(graph[v]) > 0 && dfs(u) {
			return e
		}
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	return nil
}
