package cheapestflights

import (
	"math"
	"slices"
)

// 787. K 站中转内最便宜的航班
// https://leetcode.cn/problems/cheapest-flights-within-k-stops/
// 有 n 个城市，flights[i] = (from, to, price) 表示一条单向航班。
// 求从 src 到 dst 最多经过 k 次中转（即最多乘坐 k+1 段航班）的最便宜价格，不存在返回 -1。
//
// 核心思路：带「边数限制」的最短路。
// Bellman-Ford 第 i 轮松弛恰好求出「最多 i 条边」的最短路，跑 k+1 轮即可；
// 注意每轮必须基于上一轮的 dist 副本松弛，否则一轮内可能连走多条边。

// findCheapestPrice 解法一：Bellman-Ford（限制轮数）⭐ 最优方案
// 时间复杂度 O(k · E)，空间复杂度 O(n)。
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[src] = 0
	for round := 0; round <= k; round++ {
		prev := slices.Clone(dist) // 只用上一轮结果，保证本轮最多多走一条边
		for _, f := range flights {
			u, v, w := f[0], f[1], f[2]
			if prev[u] != math.MaxInt && prev[u]+w < dist[v] {
				dist[v] = prev[u] + w
			}
		}
	}
	if dist[dst] == math.MaxInt {
		return -1
	}
	return dist[dst]
}

// findCheapestPriceBFS 解法二：按层 BFS + 剪枝（对照用）
// 第 i 层表示乘坐了 i 段航班，最多扩展 k+1 层；
// 只有到达某城市的价格比已知最低价更低时才继续入队
// （已知最低价来自更少或相同的航段数，因此剪枝是安全的）。
// 时间复杂度 O(k · E)（最坏），空间复杂度 O(n + E)。
func findCheapestPriceBFS(n int, flights [][]int, src int, dst int, k int) int {
	graph := make([][][2]int, n)
	for _, f := range flights {
		graph[f[0]] = append(graph[f[0]], [2]int{f[1], f[2]})
	}
	best := make([]int, n)
	for i := range best {
		best[i] = math.MaxInt
	}
	best[src] = 0
	q := [][2]int{{src, 0}} // [城市, 当前花费]
	for step := 0; step <= k && len(q) > 0; step++ {
		var next [][2]int
		for _, cur := range q {
			u, cost := cur[0], cur[1]
			for _, e := range graph[u] {
				v, nc := e[0], cost+e[1]
				if nc < best[v] {
					best[v] = nc
					next = append(next, [2]int{v, nc})
				}
			}
		}
		q = next
	}
	if best[dst] == math.MaxInt {
		return -1
	}
	return best[dst]
}
