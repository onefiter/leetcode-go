package networkdelay

import (
	"container/heap"
	"math"
)

// 743. 网络延迟时间
// https://leetcode.cn/problems/network-delay-time/
// 有 n 个网络节点（编号 1..n），times[i] = (u, v, w) 表示信号从 u 传到 v 需要 w 时间。
// 从节点 k 发出信号，返回所有节点都收到信号所需的最短时间；若有节点收不到返回 -1。
//
// 核心思路：单源最短路，答案是 k 到所有点最短距离的最大值。
// 边权非负，用堆优化 Dijkstra；也可以用 Bellman-Ford 做 n-1 轮松弛对照。

// distHeap 小根堆，元素为 [dist, node]。
type distHeap [][2]int

func (h distHeap) Len() int           { return len(h) }
func (h distHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h distHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *distHeap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *distHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// networkDelayTime 解法一：堆优化 Dijkstra ⭐ 最优方案
// 每次取出当前距离最小的点，它的距离已确定，再用它松弛邻居。
// 时间复杂度 O(E log E)，空间复杂度 O(n + E)。
func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][][2]int, n+1)
	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], [2]int{t[1], t[2]})
	}
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[k] = 0
	h := &distHeap{{0, k}}
	for h.Len() > 0 {
		cur := heap.Pop(h).([2]int)
		d, u := cur[0], cur[1]
		if d > dist[u] {
			continue // 过期的堆元素
		}
		for _, e := range graph[u] {
			v, w := e[0], e[1]
			if nd := d + w; nd < dist[v] {
				dist[v] = nd
				heap.Push(h, [2]int{nd, v})
			}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt {
			return -1
		}
		ans = max(ans, dist[i])
	}
	return ans
}

// networkDelayTimeBellmanFord 解法二：Bellman-Ford（对照用）
// 最短路最多 n-1 条边，对所有边松弛 n-1 轮即可；某轮无更新可提前退出。
// 时间复杂度 O(n · E)，空间复杂度 O(n)。
func networkDelayTimeBellmanFord(times [][]int, n int, k int) int {
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[k] = 0
	for round := 1; round < n; round++ {
		updated := false
		for _, t := range times {
			u, v, w := t[0], t[1], t[2]
			if dist[u] != math.MaxInt && dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				updated = true
			}
		}
		if !updated {
			break
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt {
			return -1
		}
		ans = max(ans, dist[i])
	}
	return ans
}
