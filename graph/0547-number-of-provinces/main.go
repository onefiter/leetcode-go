package provinces

import "leetcode-go/template"

// 547. 省份数量
// https://leetcode.cn/problems/number-of-provinces/
// 有 n 个城市，isConnected[i][j] = 1 表示城市 i 与 j 直接相连，
// 相连关系可传递。「省份」是一组直接或间接相连的城市，返回省份数量。
//
// 核心思路：本质是求无向图的连通分量个数。
// 并查集初始 n 个分量，每成功合并一次分量数减一；
// 也可以对每个未访问城市做一次 DFS，DFS 的启动次数就是答案。

// findCircleNum 解法一：并查集 ⭐ 最优方案
// 只遍历上三角矩阵，逐条边 Union，最后 Count 即连通分量数。
// 时间复杂度 O(n² · α(n))，空间复杂度 O(n)。
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	uf := template.NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.Count
}

// findCircleNumDFS 解法二：DFS 标记连通块（对照用）
// 从每个未访问城市出发把整个连通块染色，启动 DFS 的次数即省份数。
// 时间复杂度 O(n²)，空间复杂度 O(n)（递归栈 + visited）。
func findCircleNumDFS(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		for v := 0; v < n; v++ {
			if isConnected[u][v] == 1 && !visited[v] {
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
