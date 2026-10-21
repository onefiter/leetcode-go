package components

import "leetcode-go/template"

// 323. 无向图中连通分量的数目
// https://leetcode.cn/problems/number-of-connected-components-in-an-undirected-graph/
// 给定 n 个节点（0..n-1）和无向边列表 edges，返回图中连通分量的数目。
//
// 核心思路：并查集初始有 n 个分量，每成功合并一条边分量数 -1；
// 也可以对每个未访问节点做一次 DFS，DFS 的启动次数就是分量数。

// countComponents 解法一：并查集 ⭐ 最优方案
// 时间复杂度 O(n + E*α(n))，空间复杂度 O(n)。
func countComponents(n int, edges [][]int) int {
	uf := template.NewUnionFind(n)
	for _, e := range edges {
		uf.Union(e[0], e[1])
	}
	return uf.Count
}

// countComponentsDFS 解法二：邻接表 + DFS（对照用）
// 时间复杂度 O(n+E)，空间复杂度 O(n+E)。
func countComponentsDFS(n int, edges [][]int) int {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	visited := make([]bool, n)
	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		for _, v := range graph[u] {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			dfs(i)
		}
	}
	return count
}
