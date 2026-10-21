package genemutation

// 433. 最小基因变化
// https://leetcode.cn/problems/minimum-genetic-mutation/
// 基因序列是长度为 8、由 'A'、'C'、'G'、'T' 组成的字符串。一次变化指改变一个字符，
// 且变化后的序列必须在基因库 bank 中。返回从 startGene 变到 endGene 的最少变化次数，无法完成返回 -1。
//
// 核心思路：每个基因序列是图中的一个点，只差一个字符且在 bank 中的序列之间有边，
// 求最短路径 → BFS。枚举每一位换成其余 3 个字符，邻居最多 8*3=24 个。

// minMutation 解法一：BFS + 逐位替换枚举邻居 ⭐ 最优方案
// 时间复杂度 O(N*L*4)，空间复杂度 O(N*L)，N 为 bank 大小，L=8。
func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	valid := map[string]bool{}
	for _, b := range bank {
		valid[b] = true
	}
	if !valid[endGene] {
		return -1
	}
	visited := map[string]bool{startGene: true}
	queue := []string{startGene}
	for step := 1; len(queue) > 0; step++ {
		var next []string
		for _, g := range queue {
			bs := []byte(g)
			for i := range bs {
				orig := bs[i]
				for _, ch := range []byte("ACGT") {
					if ch == orig {
						continue
					}
					bs[i] = ch
					ng := string(bs)
					if valid[ng] && !visited[ng] {
						if ng == endGene {
							return step
						}
						visited[ng] = true
						next = append(next, ng)
					}
				}
				bs[i] = orig
			}
		}
		queue = next
	}
	return -1
}

// minMutationGraph 解法二：预处理 bank 两两建图再 BFS（对照用）
// 把 startGene 和 bank 中所有序列作为点，两两比较是否只差一个字符来建邻接表，再跑 BFS。
// 时间复杂度 O(N^2*L)，空间复杂度 O(N^2)。
func minMutationGraph(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	nodes := append([]string{startGene}, bank...)
	diffOne := func(a, b string) bool {
		diff := 0
		for i := range a {
			if a[i] != b[i] {
				diff++
			}
		}
		return diff == 1
	}
	n := len(nodes)
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if diffOne(nodes[i], nodes[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}
	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	dist[0] = 0
	queue := []int{0}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		if nodes[u] == endGene {
			return dist[u]
		}
		for _, v := range graph[u] {
			if dist[v] == -1 {
				dist[v] = dist[u] + 1
				queue = append(queue, v)
			}
		}
	}
	return -1
}
