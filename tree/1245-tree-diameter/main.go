package treediameter

// 1245. 树的直径
// https://leetcode.cn/problems/tree-diameter/
// 给一棵无向树（n 个节点编号 0..n-1，用 n-1 条边 edges 描述），
// 返回树的直径：树中最长路径包含的边数。
//
// 核心思路：
//   - 两次 BFS：从任意点出发找到最远点 a，a 一定是某条直径的端点；
//     再从 a 出发找到最远点，距离就是直径。
//   - 树形 DP：以任意点为根 DFS，对每个节点求「向下最长链」；
//     经过该节点的最长路径 = 最长的两条子链之和，取全局最大。

// treeDiameter 解法一：两次 BFS ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(n)。
func treeDiameter(edges [][]int) int {
	n := len(edges) + 1
	g := buildGraph(n, edges)
	a, _ := farthest(g, 0)
	_, d := farthest(g, a)
	return d
}

func buildGraph(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	return g
}

// farthest 从 src 出发 BFS，返回最远的节点及其距离。
func farthest(g [][]int, src int) (node, dist int) {
	d := make([]int, len(g))
	for i := range d {
		d[i] = -1
	}
	d[src] = 0
	queue := []int{src}
	node = src
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		if d[u] > d[node] {
			node = u
		}
		for _, v := range g[u] {
			if d[v] == -1 {
				d[v] = d[u] + 1
				queue = append(queue, v)
			}
		}
	}
	return node, d[node]
}

// treeDiameterDP 解法二：树形 DP（DFS 求最长两条子链）
// down(u) = 从 u 往子树方向走的最长边数；经过 u 的最长路径 = 最长子链 + 次长子链。
// 时间复杂度 O(n)，空间复杂度 O(n)（递归栈最坏 O(n)）。
func treeDiameterDP(edges [][]int) int {
	n := len(edges) + 1
	g := buildGraph(n, edges)
	best := 0
	var dfs func(u, parent int) int
	dfs = func(u, parent int) int {
		first, second := 0, 0
		for _, v := range g[u] {
			if v == parent {
				continue
			}
			l := dfs(v, u) + 1
			if l > first {
				first, second = l, first
			} else if l > second {
				second = l
			}
		}
		best = max(best, first+second)
		return first
	}
	dfs(0, -1)
	return best
}

// treeDiameterEnumeration 解法三：从每个节点 BFS 求最远距离（对照用）
// 直径 = 所有节点对之间距离的最大值。
// 时间复杂度 O(n^2)，空间复杂度 O(n)。
func treeDiameterEnumeration(edges [][]int) int {
	n := len(edges) + 1
	g := buildGraph(n, edges)
	best := 0
	for s := 0; s < n; s++ {
		_, d := farthest(g, s)
		best = max(best, d)
	}
	return best
}
