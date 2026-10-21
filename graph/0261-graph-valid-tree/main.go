package validtree

import "leetcode-go/template"

// 261. 以图判树
// https://leetcode.cn/problems/graph-valid-tree/
// 给定 n 个节点（0..n-1）和无向边列表 edges，判断这些边是否能构成一棵有效的树。
//
// 核心思路：n 个节点的树 ⇔ 恰好 n-1 条边且连通（等价于 n-1 条边且无环）。
// 先用边数剪枝，再用并查集判环 / 或 BFS 判断是否连通。

// validTree 解法一：边数检查 + 并查集判环 ⭐ 最优方案
// 时间复杂度 O(n*α(n))，空间复杂度 O(n)。
func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}
	uf := template.NewUnionFind(n)
	for _, e := range edges {
		// 两端已连通再加边就成环
		if !uf.Union(e[0], e[1]) {
			return false
		}
	}
	return true
}

// validTreeBFS 解法二：边数检查 + BFS 判连通（对照用）
// 边数为 n-1 时，只要从 0 出发能到达所有节点就一定无环。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func validTreeBFS(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	for _, d := range template.GraphBFS(graph, 0) {
		if d == -1 {
			return false
		}
	}
	return true
}
